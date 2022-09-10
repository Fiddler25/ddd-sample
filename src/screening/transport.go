package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/sdk/encode"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(ctx context.Context, client *ent.Client) http.Handler {
	screeningRepo := NewScreeningRepository(ctx, client)

	startFromPreInterviewHandler := kithttp.NewServer(
		makeStartFromPreInterview(NewScreeningUseCase(screeningRepo)),
		decodeStartFromPreInterviewRequest,
		encode.Response,
	)
	applyHandler := kithttp.NewServer(
		makeApply(NewScreeningUseCase(screeningRepo)),
		decodeApplyRequest,
		encode.Response,
	)
	addNextInterviewHandler := kithttp.NewServer(
		makeAddNextInterview(NewScreeningUseCase(screeningRepo)),
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
