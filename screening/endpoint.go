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
		ID                    ID     `json:"id"`
		Status                Status `json:"status"`
		ApplicantEmailAddress string `json:"applicant_email_address"`
		Err                   error  `json:"err,omitempty"`
	}
)

func (r startFromPreInterviewResponse) error() error { return r.Err }

func makeStartFromPreInterviewEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(startFromPreInterviewRequest)
		out, err := s.StartFromPreInterview(ctx, req.ApplicantEmailAddress)

		return startFromPreInterviewResponse{
			ID:                    out.ID,
			Status:                out.Status,
			ApplicantEmailAddress: out.ApplicantEmailAddress,
			Err:                   err,
		}, nil
	}
}
