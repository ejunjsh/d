package network

import "net"

type Network struct {
	Name    string
	IpRange *net.IPNet
	Driver  string
}
