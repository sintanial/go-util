package httputil

import (
	"net/http"
	"crypto/tls"
	"net/http/httputil"
	"net"
	"bitbucket.org/MountAim/go-util/netutil"
	"bytes"
	"io/ioutil"
)

func NewResponse(statusCode int, status string, headers http.Header) *http.Response {
	res := &http.Response{
		Status:     status,
		StatusCode: statusCode,
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	if headers != nil {
		res.Header = headers
	}

	return res
}

func NewResponseBody(statusCode int, status string, headers http.Header, body []byte) *http.Response {
	res := NewResponse(statusCode, status, headers)

	if body != nil {
		buf := bytes.NewBuffer(body)
		res.Body = ioutil.NopCloser(buf)
	}

	return res
}

func Dump(r interface{}, body bool) string {
	if r == nil {
		return ""
	}

	var data []byte
	var err error
	switch re := r.(type) {
	case *http.Request:
		if re == nil {
			return ""
		}

		data, err = httputil.DumpRequestOut(re, body)
		break
	case *http.Response:
		if re == nil {
			return ""
		}

		data, err = httputil.DumpResponse(re, body)
		break
	default:
		return ""
	}

	if err != nil {
		return ""
	}

	return string(data)
}

func ListenTLS(server *http.Server, tlsconf *tls.Config) error {
	listener, err := tls.Listen("tcp", server.Addr, tlsconf)
	if err != nil {
		return err
	}

	return server.Serve(listener)
}

func ParseClientIP(r *http.Request) net.IP {
	if r.Header.Get("Client-Ip") != "" {
		return netutil.ParseIP(r.Header.Get("Client-Ip"))
	} else if r.Header.Get("X-Forwarded-For") != "" {
		return netutil.ParseIP(r.Header.Get("X-Forwarded-For"))
	} else if r.Header.Get("X-Real-Ip") != "" {
		return netutil.ParseIP(r.Header.Get("X-Real-Ip"))
	} else {
		return netutil.ParseIP(r.RemoteAddr)
	}
}

func RequestScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}

	return "http"
}

type ServeMux struct {
	muxs map[string]http.Handler
}

func NewServeMux() *ServeMux {
	return new(ServeMux)
}

func (self *ServeMux) make() {
	if self.muxs == nil {
		self.muxs = make(map[string]http.Handler)
	}
}

func (self *ServeMux) Handle(pattern string, handler http.Handler) {
	self.make()

	self.muxs[pattern] = handler
}

// HandleFunc registers the handler function for the given pattern.
func (self *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	self.Handle(pattern, http.HandlerFunc(handler))
}

func (self *ServeMux) Add(mux *ServeMux) {
	self.make()

	for p, h := range mux.muxs {
		if _, ok := self.muxs[p]; ok {
			panic("http: multiple registrations for " + p)
		}

		self.muxs[p] = h
	}
}

func (self *ServeMux) HttpMux() *http.ServeMux {
	m := http.NewServeMux()
	self.AppendTo(m)
	return m
}

type Muxator interface {
	Handle(pattern string, handler http.Handler)
}

func (self *ServeMux) AppendTo(m Muxator) {
	for p, h := range self.muxs {
		m.Handle(p, h)
	}
}

func MuxCombine(mux ...*ServeMux) *ServeMux {
	res := &ServeMux{}
	for _, m := range mux {
		res.Add(m)
	}

	return res
}
