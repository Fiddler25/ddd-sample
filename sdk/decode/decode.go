package decode

import (
	"context"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/exp/constraints"
	"net/http"
	"strconv"
)

type Request[T constraints.Integer] struct {
	ID T
}

func RequestID[T constraints.Integer](path string) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		id, err := FromID[T](r, path)
		if err != nil {
			return nil, err
		}
		return Request[T]{ID: id}, nil
	}
}

func FromID[T constraints.Integer](r *http.Request, path string) (T, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[path])
	if err != nil || id <= 0 {
		return 0, errors.New("invalid request")
	}
	return T(id), nil
}
