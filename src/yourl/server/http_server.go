package server

import (
	"fmt"
	"net/http"
	"bytes"
)

/**
 * Http Server
 */
type HttpServer struct {
	srv *http.ServeMux
	domain string
	version string
	urlMethod string
	port string
}

func (h *HttpServer) getUrlPath() string {
	return "/" + h.version + "/" + h.urlMethod
}

func (h *HttpServer) getAddr() string {
	return h.domain + ":" + h.port
}

func (h *HttpServer) Run() {
	h.version = "v1"
	h.urlMethod = "url"
	h.port = "4000"
	h.srv = http.NewServeMux()
	// API Path
	h.srv.HandleFunc(h.getUrlPath(), apiHandler)
	// URL
	h.srv.HandleFunc("/", resolveHandler)
	http.ListenAndServe(h.getAddr(), h.srv)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API - Do %s!\nMétodo: %s\n", r.URL.Path[1:], r.Method	)
	if r.Method == "POST" {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		fmt.Fprintf(w, "Body = %s\n", buf.String())
	}

}
func resolveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Redirect to %s!", r.URL.Path[1:])
}
