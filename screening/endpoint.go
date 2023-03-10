package screening

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type (
	StartFromPreInterviewRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address"`
	}

	StartFromPreInterviewResponse struct {
		ID                    ID     `json:"id,omitempty"`
		Status                Status `json:"status,omitempty"`
		ApplicantEmailAddress string `json:"applicant_email_address,omitempty"`
		Err                   error  `json:"err,omitempty"`
	}
)

func (r StartFromPreInterviewResponse) error() error { return r.Err }

func makeStartFromPreInterviewEndpoint(ctx context.Context, s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(StartFromPreInterviewRequest)
		out, err := s.StartFromPreInterview(ctx, req.ApplicantEmailAddress)

		return &StartFromPreInterviewResponse{
			ID:                    out.ID,
			Status:                out.Status,
			ApplicantEmailAddress: out.ApplicantEmailAddress,
			Err:                   err,
		}, nil
	}
}
