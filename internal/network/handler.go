package network

import (
	. "docker-/internal/log"
	"docker-/pkg/utils"
	"net"
)

func Create(model *CreateModel) error {

	// parse subnet CIDR
	_, ipNet, err := net.ParseCIDR(model.Subnet)
	if err != nil {
		return err
	}

	ipam := &IPAM{
		FilePath: "",
		BitMap:   make([]byte, GetIPCount(ipNet)),
	}

	gateway := GetGateway(ipNet)

	network := &Network{
		Id:         "",
		Name:       model.Name,
		Driver:     model.Driver,
		Refers:     0,
		IpNet:      ipNet,
		Gateway:    gateway,
		Ipam:       ipam,
		CreateTime: utils.TimeNowToString(),
	}

	Log.Debugf("network info: %v", network)

	return nil
}

func Remove(identifier string) error {

	return nil
}

func List() error {
	return nil
}

func checkInterface(subnet string) error {

	return nil
}
