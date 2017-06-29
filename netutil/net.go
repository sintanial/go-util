package netutil

import (
	"net"
	"strings"
)

func ParseIP(addr string) net.IP {
	host := strings.Split(addr, ":")[0]
	if host == "" {
		return nil
	}

	return net.ParseIP(host)
}

func MustParseCIDR(s string) *net.IPNet {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}

	return ipnet
}