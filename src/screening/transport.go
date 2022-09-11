package screening

import (
	"context"
	"ddd-sample/sdk/encode"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(uc Usecase) http.Handler {
	startFromPreInterviewHandler := kithttp.NewServer(
		makeStartFromPreInterview(uc),
		decodeStartFromPreInterviewRequest,
		encode.Response,
	)
	applyHandler := kithttp.NewServer(
		makeApply(uc),
		decodeApplyRequest,
		encode.Response,
	)
	addNextInterviewHandler := kithttp.NewServer(
		makeAddNextInterview(uc),
		decodeAddNextInterview,
		encode.Response,
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
