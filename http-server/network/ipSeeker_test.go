package network

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

type dummyAddr string

func (a dummyAddr) Network() string {
	return string(a)
}
func (a dummyAddr) String() string {
	return string(a)
}

func TestGetLocalIpReturnErrorWhenInterfaceAddrsReturnError(t *testing.T) {
	var stub = func() ([]net.Addr, error) {
		return nil, fmt.Errorf("error during getting interface adresses")
	}
	_, err := GetLocalIp(stub)
	assert.Error(t, err)
	assert.Equal(t, "error during getting interface adresses", err.Error())
}

func TestGetLocalIpReturnErrorWhenInterfaceAddrsReturnNothing(t *testing.T) {
	var stub = func() ([]net.Addr, error) {
		return nil, nil
	}
	_, err := GetLocalIp(stub)
	assert.Error(t, err)
	assert.Equal(t, "no local ip found", err.Error())
}

func TestGetLocalIpReturnLocalIpWhenInterfaceAddrsReturnGoodIpNet(t *testing.T) {
	var stub = func() ([]net.Addr, error) {
		goodAddr := &net.IPNet{
			IP: net.ParseIP("128.168.0.44"),
		}
		return []net.Addr{goodAddr}, nil
	}
	ip, err := GetLocalIp(stub)
	assert.Equal(t, "128.168.0.44", ip.String())
	assert.NoError(t, err)
}

func TestGetLocalIpReturnErrorWhenInterfaceAddrsReturnLoopbackIPNetAddrs(t *testing.T) {
	var stub = func() ([]net.Addr, error) {
		return []net.Addr{dummyAddr("not good address")}, nil
	}
	_, err := GetLocalIp(stub)
	assert.Error(t, err)
	assert.Equal(t, "no local ip found", err.Error())
}

func TestGetLocalIpReturnErrorWhenInterfaceAddrsReturnNoIPNetAddrs(t *testing.T) {
	var stub = func() ([]net.Addr, error) {
		loopbackIp := &net.IPNet{
			IP: net.ParseIP("127.0.0.1"),
		}
		return []net.Addr{loopbackIp}, nil
	}
	_, err := GetLocalIp(stub)
	assert.Error(t, err)
	assert.Equal(t, "no local ip found", err.Error())
}

func TestGetLocalIpReturnErrorWhenInterfaceAddrsReturnIPv6(t *testing.T) {
	var stub = func() ([]net.Addr, error) {
		localIpv6 := &net.IPNet{
			IP: net.ParseIP("ff02::1"),
		}
		return []net.Addr{localIpv6}, nil
	}
	_, err := GetLocalIp(stub)
	assert.Error(t, err)
	assert.Equal(t, "no local ip found", err.Error())
}
