package network

import (
	"github.com/vishvananda/netlink"
	"net"
)

type Network struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Driver     string     `json:"driver"`
	Refers     uint64     `json:"refer"`
	IpNet      *net.IPNet `json:"ip_net"`
	Gateway    *net.IP    `json:"gateway"`
	Ipam       *IPAM      `json:"ipam"`
	CreateTime string     `json:"create_time"`
}

type Endpoint struct {
	Id          string            `json:"id"`
	Network     *Network          `json:"network"`
	Device      *netlink.Device   `json:"device"`
	IPAddr      *net.IPAddr       `json:"ip_addr"`
	MacAddr     *net.HardwareAddr `json:"mac_addr"`
	PortMapping []string          `json:"port_hash"`
}

type Driver interface {
	Name() string
	Create() (*Network, error)
	Remove() error
	Connect() error
	DisConnect() error
}

type IPAM struct {
	FilePath string `json:"file_path"`
	BitMap   []byte `json:"bit_map"`
}

////////////////////////////////////// Params ///////////////////////////////////////////////

type CreateModel struct {
	Name   string `json:"name"`
	Driver string `json:"driver"`
	Subnet string `json:"subnet"`
}
