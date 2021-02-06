package main

import (
	"net/http"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ip_addr := r.Header.Get("X-Real-Ip")
  if ip_addr == "" {
		ip_addr = r.Header.Get("X-Forwarded-For")
	}
  if ip_addr == "" {
		ip_addr = r.RemoteAddr
  }

	w.Write([]byte(ip_addr + "\n"))
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ipHandler)
	http.ListenAndServe(":80", mux)
}
