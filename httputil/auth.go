package httputil

import (
	"strings"
	"encoding/base64"
	"net/http"
)

type HttpAuthCredential struct {
	Username string
	Password string
}

type HttpAuth struct {
	Creds []*HttpAuthCredential
	Realm string
}

func NewHttpAuth(creds []*HttpAuthCredential, realm string) *HttpAuth {
	return &HttpAuth{creds, realm}
}

func (self *HttpAuth) IsValidAuthRequest(r *http.Request) bool {
	username, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	if len(self.Creds) == 0 {
		return true
	}

	for _, cred := range self.Creds {
		if cred.Username == username && cred.Password == pass {
			return true
		}
	}

	return false
}

func (self *HttpAuth) AuthHandler(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", `*`)

		if r.Method == "OPTIONS" {
			reqm := r.Header.Get("Access-Control-Request-Method")
			if reqm != "" {
				w.Header().Set("Access-Control-Allow-Method", reqm)
			}

			reqh := r.Header.Get("Access-Control-Request-Headers")
			if reqh != "" {
				w.Header().Set("Access-Control-Allow-Headers", reqh)
			}

			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		if !self.Auth(w, r) {
			return
		}

		handler(w, r)
	}
}

func (self *HttpAuth) Auth(w http.ResponseWriter, r *http.Request) bool {
	if !self.IsValidAuthRequest(r) {
		w.Header().Set("WWW-Authenticate", `Basic realm="` + self.Realm + `"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func ParseBasicAuth(auth string) (username, password string, ok bool) {
	if !strings.HasPrefix(auth, "Basic ") {
		return
	}

	c, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}

	return cs[:s], cs[s + 1:], true
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}