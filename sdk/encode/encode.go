package encode

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type errorer interface {
	error() error
}

func Response(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err, ok := response.(errorer); ok && err.error() != nil {
		encodeError(err.error(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(err error, w http.ResponseWriter) {
	w.WriteHeader(codeFrom(err))
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInconsistentIDs, ErrAlreadyExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
