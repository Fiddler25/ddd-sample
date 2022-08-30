package screening

import (
	"context"
	"ddd-sample/src/usecase/screening"
	"github.com/go-kit/kit/endpoint"
)

type (
	startFromPreInterviewRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address"`
	}

	startFromPreInterviewResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func (r startFromPreInterviewResponse) error() error { return r.Err }

func makeStartFromPreInterview(uc screening.ScreeningUseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(startFromPreInterviewRequest)
		err := uc.StartFromPreInterview(req.ApplicantEmailAddress)
		return startFromPreInterviewResponse{Err: err}, nil
	}
}
