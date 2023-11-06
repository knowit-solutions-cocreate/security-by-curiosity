package info_test

import (
	"net"
	"testing"

	info "github.com/knowit-solutions-cocreate/security-exercise/network-info"
)

func TestIsIpGetInternalIp(t *testing.T) {
	ip, err := info.GetInternalIp()
	if err == nil {
		_, err = net.ResolveIPAddr("ip", ip)
	} else {
		t.Errorf("not an ip: %s", ip)
	}

}
func TestExternalIsNotRFC1918(t *testing.T) {
	ip_ranges := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}
	ip, err := info.GetExternalIp()
	if err == nil {
		for _, rfc1918_ip_range := range ip_ranges {
			_, ipNet, _ := net.ParseCIDR(rfc1918_ip_range)
			ipToCheck := net.ParseIP(ip)
			if ipNet.Contains(ipToCheck) {
				t.Errorf("%s resides in %s, :(", ip, rfc1918_ip_range)
			}
		}
	} else {
		t.Errorf("external ip is rfc1918: %s", ip)
	}

}
