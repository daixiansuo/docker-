package network

import (
	"fmt"
	"net"
	"testing"
)

func TestParseCIDR(t *testing.T) {
	ip, IpNet, _ := net.ParseCIDR("10.10.0.6/24")

	fmt.Printf("ip: %v, IpNet: %v", ip, IpNet)
}

func TestNetInterface(t *testing.T) {

	_, ipNet, _ := net.ParseCIDR("10.10.0.6/24")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ar := addr.String()
		if ar == "192.168.0.1" {
			fmt.Errorf("the subnet %s already exists", ipNet)
			return
		}
	}
}

func TestGetIPCount(t *testing.T) {
	_, ipNet, _ := net.ParseCIDR("10.10.0.6/24")

	println(GetIPCount(ipNet))
}
