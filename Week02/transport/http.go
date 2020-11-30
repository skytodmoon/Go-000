package transport

import (
	"context"
	"encoding/json"

	//"errors"
	"fmt"
	"net/http"
	"os"

	"Week02/endpoint"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func MakeHTTPHandler(ctx context.Context, endpoints *endpoint.UserEndpoints) http.Handler {
	r := mux.NewRouter()

	kitLog := log.NewLogfmtLogger(os.Stderr)

	kitLog = log.With(kitLog, "ts", log.DefaultTimestampUTC)
	kitLog = log.With(kitLog, "caller", log.DefaultCaller)

	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods("GET").Path("/queryuser").Handler(kithttp.NewServer(
		endpoints.QueryUserEndpoint,
		decodeQueryUserRequest,
		encodeJSONResponse,
		options...,
	))
	return r
}

func decodeQueryUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	email := r.FormValue("email")
	fmt.Println("email:", email)

	if email == "" {
		return nil, ErrorBadRequest
	}
	return &endpoint.QueryUserRequest{
		Email: email,
	}, nil
	return nil, ErrorBadRequest
}

func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
	fmt.Printf("Error:%+v\n", err)
}
