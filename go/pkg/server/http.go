package server

import (
	"net/http"

	"dbut.sh/pkg/provider"
)

type HTTPServer struct {
	Address  string
	Provider provider.ContentProvider
}

func NewHTTPServer(address string, p provider.ContentProvider) *HTTPServer {
	return &HTTPServer{
		Address:  address,
		Provider: p,
	}
}

func (s *HTTPServer) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", s.Provider.GetContentType())
		if err := provider.WriteContent(w, s.Provider, r.URL.Path); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return http.ListenAndServe(s.Address, nil)
}

func StartHTTPServer(address string, p provider.ContentProvider) error {
	s := NewHTTPServer(address, p)
	return s.Start()
}
