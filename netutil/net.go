package netutil

import (
	"net"
	"strings"
	"strconv"
	"encoding/binary"
)

func ParseIP(addr string) net.IP {
	addr, _ = SplitHostPort(addr)
	return net.ParseIP(addr)
}

func SplitHostPort(s string) (string, int) {
	idx := strings.LastIndex(s, ":")
	var host string
	var port int
	if idx >= 0 {
		host = s[:idx]
		port, _ = strconv.Atoi(s[idx:])
	} else {
		host = s
	}

	return host, port
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
