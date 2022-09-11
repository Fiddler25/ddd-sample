package screening

import (
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

func MakeHandler(ctx context.Context, uc Usecase) http.Handler {
	startFromPreInterviewHandler := kithttp.NewServer(
		makeStartFromPreInterview(ctx, uc),
		decodeStartFromPreInterviewRequest,
		encodeResponse,
	)
	applyHandler := kithttp.NewServer(
		makeApply(ctx, uc),
		decodeApplyRequest,
		encodeResponse,
	)
	addNextInterviewHandler := kithttp.NewServer(
		makeAddNextInterview(ctx, uc),
		decodeAddNextInterview,
		encodeResponse,
	)

	r := mux.NewRouter()

	r.Handle("/screening/v1/start_from_pre_interview", startFromPreInterviewHandler).Methods("POST")
	r.Handle("/screening/v1/apply", applyHandler).Methods("POST")
	r.Handle("/screening/v1/add_next_interview", addNextInterviewHandler).Methods("PUT")

	return r
}

func decodeStartFromPreInterviewRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req startFromPreInterviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeApplyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req applyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAddNextInterview(_ context.Context, r *http.Request) (interface{}, error) {
	var req addNextInterviewRequest
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrInconsistentIDs, ErrAlreadyExists:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
