package screening

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	startFromPreInterviewRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address"`
	}

	startFromPreInterviewResponse struct {
		Err error `json:"err,omitempty"`
	}
)

func (r startFromPreInterviewResponse) error() error { return r.Err }

func makeStartFromPreInterviewEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(startFromPreInterviewRequest)
		err := s.StartFromPreInterview(ctx, req.ApplicantEmailAddress)

		return startFromPreInterviewResponse{Err: err}, nil
	}
}
