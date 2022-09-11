package screening

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"time"
)

// StartFromPreInterview
type (
	startFromPreInterviewRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address"`
	}

	startFromPreInterviewResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func (r startFromPreInterviewResponse) error() error { return r.Err }

func makeStartFromPreInterview(ctx context.Context, uc Usecase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(startFromPreInterviewRequest)
		err := uc.StartFromPreInterview(ctx, req.ApplicantEmailAddress)
		return startFromPreInterviewResponse{Err: err}, nil
	}
}

// Apply
type (
	applyRequest struct {
		ApplicantEmailAddress string `json:"applicant_email_address"`
	}

	applyResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func (r applyResponse) error() error { return r.Err }

func makeApply(ctx context.Context, uc Usecase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(applyRequest)
		err := uc.Apply(ctx, req.ApplicantEmailAddress)
		return applyResponse{Err: err}, nil
	}
}

// AddNextInterview
type (
	addNextInterviewRequest struct {
		ScreeningID   string    `json:"screening_id"`
		InterviewDate time.Time `json:"interview_date"`
	}

	addNextInterviewResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func (r addNextInterviewResponse) error() error { return r.Err }

func makeAddNextInterview(ctx context.Context, uc Usecase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addNextInterviewRequest)
		err := uc.AddNextInterview(ctx, req.ScreeningID, req.InterviewDate)
		return addNextInterviewResponse{Err: err}, nil
	}
}
