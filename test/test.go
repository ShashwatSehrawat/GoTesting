package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
)

func TestMockConsulServer(t *testing.T) {
	// Create a new instance of the Consul API client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		t.Fatalf("Failed to create Consul client: %s", err)
	}

	// Create a new HTTP server using httptest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle service registration and deregistration requests
		switch r.URL.Path {
		case "/v1/agent/service/register":
			// Process service registration request
			// ...

			// For example, assert the service name and check if it's registered correctly
			assert.Equal(t, "my-service", r.FormValue("Name"))
			assert.Equal(t, "localhost:8080", r.FormValue("Address"))
			assert.Equal(t, "http", r.FormValue("Check.HTTP"))
			assert.Equal(t, "/health", r.FormValue("Check.HTTP"))

			// Return a mock Consul response indicating successful registration
			fmt.Fprint(w, `{"ID": "service-id", "Name": "my-service", "Address": "localhost", "Port": 8080}`)

		case "/v1/agent/service/deregister":
			// Process service deregistration request
			// ...

			// For example, assert the service ID and check if it's deregistered correctly
			assert.Equal(t, "service-id", r.FormValue("service_id"))

			// Return a mock Consul response indicating successful deregistration
			fmt.Fprint(w, `{"ID": "service-id", "Name": "my-service", "Address": "localhost", "Port": 8080}`)
		}
	}))

	// Override Consul client's HTTP address with the mock server's address
	client.Address = server.URL

	// Perform the service registration test scenario
	registerService(t, client)

	// Perform the service deregistration test scenario
	deregisterService(t, client)

	// Close the mock server after the tests
	server.Close()
}

func registerService(t *testing.T, client *api.Client) {
	// Create a new service registration
	reg := &api.AgentServiceRegistration{
		ID:      "service-id",
		Name:    "my-service",
		Address: "localhost",
		Port:    8080,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://localhost:8080/health",
			Interval: "10s",
		},
	}

	// Register the service with Consul
	err := client.Agent().ServiceRegister(reg)
	assert.NoError(t, err)
}

func deregisterService(t *testing.T, client *api.Client) {
	// Deregister the service from Consul
	err := client.Agent().ServiceDeregister("service-id")
	assert.NoError(t, err)
}
