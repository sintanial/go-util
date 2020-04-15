package httputil

import "net/http"

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

type MiddlewareFunc func(handler http.Handler) http.Handler

func (self *ServeMux) Middleware(pattern string, fn MiddlewareFunc) bool {
	if self.muxs == nil {
		return false
	}

	original, ok := self.muxs[pattern]
	if !ok {
		return false
	}

	self.muxs[pattern] = fn(original)
	return true
}

func (self *ServeMux) MiddlewareAll(fn MiddlewareFunc) {
	for pattern, handler := range self.muxs {
		self.muxs[pattern] = fn(handler)
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
