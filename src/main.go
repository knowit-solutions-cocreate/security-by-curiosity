package main

import (
	"fmt"

	info "github.com/knowit-solutions-cocreate/security-exercise/network-info"
)

func main() {
	internal_ip, _ := info.GetInternalIp()
	external_ip, _ := info.GetExternalIp()
	fmt.Println("Internal ip address: " + internal_ip)
	fmt.Println("External ip address: " + external_ip)
}
