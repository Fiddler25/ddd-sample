package screening

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

var ErrValidation = errors.New("validation error")

type (
	StartFromPreInterviewRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address" validate:"required,email"`
	}

	StartFromPreInterviewResponse struct {
		ID                    ID     `json:"id,omitempty"`
		Status                Status `json:"status,omitempty"`
		ApplicantEmailAddress string `json:"applicant_email_address,omitempty"`
		Err                   error  `json:"err,omitempty"`
	}
)

func (r StartFromPreInterviewResponse) error() error { return r.Err }

func MakeStartFromPreInterviewEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StartFromPreInterviewRequest)

		if err := validator.New().Struct(req); err != nil {
			return StartFromPreInterviewResponse{Err: ErrValidation}, nil
		}
		out, err := s.StartFromPreInterview(ctx, req.ApplicantEmailAddress)

		return StartFromPreInterviewResponse{
			ID:                    out.ID,
			Status:                out.Status,
			ApplicantEmailAddress: out.ApplicantEmailAddress,
			Err:                   err,
		}, nil
	}
}
