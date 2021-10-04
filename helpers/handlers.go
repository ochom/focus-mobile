package helpers

import "net/http"

func Register() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func AddPrescriptions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}
