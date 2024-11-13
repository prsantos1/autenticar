package server

import (
	"net/http"
)

// Start inicia o servidor HTTP
func Start() error {
	http.HandleFunc("/login", loginHandler)
	return http.ListenAndServe(":808", nil)
}
