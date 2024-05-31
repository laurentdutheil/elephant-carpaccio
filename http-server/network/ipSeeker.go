package network

import (
	"fmt"
	"net"
)

type InterfaceAddrs = func() ([]net.Addr, error)

// GetLocalIp retrieve a local ip v4
// use: GetLocalIp(net.InterfaceAddrs) in production
// you can stub InterfaceAddrs for tests
func GetLocalIp(interfaceAddrsSeeker InterfaceAddrs) (net.IP, error) {
	addrs, err := interfaceAddrsSeeker()
	if err != nil {
		return nil, fmt.Errorf("no local ip found")
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP, nil
		}
	}
	return nil, fmt.Errorf("no local ip found")
}
