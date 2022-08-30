package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/sdk/encode"
	is "ddd-sample/src/infra/screening"
	us "ddd-sample/src/usecase/screening"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(ctx context.Context, client *ent.Client) http.Handler {
	screeningRepo := is.NewScreeningRepository(ctx, client)

	applyHandler := kithttp.NewServer(
		makeApply(us.NewScreeningUseCase(screeningRepo)),
		decodeApplyRequest,
		encode.Response,
	)

	r := mux.NewRouter()

	r.Handle("/screening/v1/apply_interview", applyHandler).Methods("POST")

	return r
}

func decodeApplyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req applyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
