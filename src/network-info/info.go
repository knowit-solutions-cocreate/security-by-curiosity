package info

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func GetExternalIp() (string, error) {
	response, err := http.Get("https://ident.me")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	external_ip, err := io.ReadAll(response.Body)

	if err != nil {
		return "", nil
	} else {
		return string(external_ip), nil
	}
}

func GetInternalIp() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue // Skip down and loopback interfaces
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue // Not an IP address
			}
			if ipnet.IP.To4() == nil {
				continue // Skip IPv6 addresses
			}
			return ipnet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("No default gateway found")

}
