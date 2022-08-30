package screening

import (
	"context"
	"ddd-sample/src/usecase/screening"
	"github.com/go-kit/kit/endpoint"
)

type (
	applyRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address"`
	}

	applyResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func (r applyResponse) error() error { return r.Err }

func makeApply(uc screening.ScreeningUseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(applyRequest)
		err := uc.Apply(req.ApplicantEmailAddress)
		return applyResponse{Err: err}, nil
	}
}