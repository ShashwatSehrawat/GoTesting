package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Name == "eth0" {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip.To4() == nil {
					fmt.Println(ip)
				}
			}
		}
	}
}
