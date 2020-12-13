package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func init() {
	decoder.IgnoreUnknownKeys(true)
}

func DecodePingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
