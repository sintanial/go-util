package httputil

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/sintanial/go-util/netutil"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
)

func IsJsonContentType(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

func ReadJson(data interface{}, r *http.Request) error {
	if !IsJsonContentType(r) {
		return errors.New("invalid content-type, must be json")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func WriteJson(data interface{}, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	_, ok := r.URL.Query()["pretty"]
	return WriteJsonTo(data, w, ok)
}

func WriteJsonTo(data interface{}, w io.Writer, isPretty bool) error {
	var jdata []byte
	var err error

	if isPretty {
		jdata, err = json.MarshalIndent(data, "", "  ")
	} else {
		jdata, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	_, err = w.Write(jdata)
	return err
}

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