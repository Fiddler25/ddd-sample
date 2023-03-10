package screening

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(ctx context.Context, s Service) http.Handler {
	r := mux.NewRouter()

	startFromPreInterviewHandler := kithttp.NewServer(
		makeStartFromPreInterviewEndpoint(ctx, s),
		decodeStartFromPreInterviewRequest,
		encodeResponse,
	)
	r.Methods("POST").Path("/screening/v1/start-from-pre-interview").Handler(startFromPreInterviewHandler)

	return r
}

func decodeStartFromPreInterviewRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req startFromPreInterviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
