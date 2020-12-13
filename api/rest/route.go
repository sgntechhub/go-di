package rest

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func newHTTPHandler(endpoint Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/ping").Handler(httptransport.NewServer(
		endpoint.Ping,
		DecodePingRequest,
		EncodeHTTPResponse,
	))
	return r
}
