package netutil

import (
	"net"
	"strings"
	"encoding/binary"
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

func IpToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func IntToIp(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

func IpRangeToCIDR(ipFrom net.IP, ipTo net.IP) *net.IPNet {
	max := 32
	for l := max; l >= 0; l-- {
		mask := net.CIDRMask(l, max)

		na := ipFrom.Mask(mask)
		n := net.IPNet{IP: na, Mask: mask}

		if n.Contains(ipTo) {
			return &n
		}
	}

	return nil
}
