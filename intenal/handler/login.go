package handler

import "net/http"

// Login ...
func Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{'s'})
}
