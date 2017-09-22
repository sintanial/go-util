package urlutil

import (
	"strings"
	"net/url"
)

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

func MustParseUrl(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic("failed to parse rawurl:" + err.Error())
	}

	return u
}
