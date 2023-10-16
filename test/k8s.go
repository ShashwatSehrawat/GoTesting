// package main

// import (
// 	"context"
// 	"fmt"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// func main() {
// 	// Load kubeconfig
// 	config, err := clientcmd.BuildConfigFromFlags("", "~/.kube/config")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create Kubernetes client
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get pod by name and namespace
// 	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "my-pod", metav1.GetOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get pod IP address
// 	podIP := pod.Status.PodIP

// 	fmt.Println("Pod IP:", podIP)
// }
