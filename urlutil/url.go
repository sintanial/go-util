package urlutil

import "strings"

func SplitHostPort(s string) (host string, port string) {
	split := strings.Split(s, ":")
	host = split[0]
	if len(split) > 1 {
		port = split[1];
	}

	return host, port
}

func NormalizeHostPort(host string) string {
	h, p := SplitHostPort(host)
	if p == "80" || p == "443" {
		return h
	}

	return host
}
