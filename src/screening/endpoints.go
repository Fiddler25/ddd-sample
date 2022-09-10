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

func makeStartFromPreInterview(uc ScreeningUseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(startFromPreInterviewRequest)
		err := uc.StartFromPreInterview(req.ApplicantEmailAddress)
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

func makeApply(uc ScreeningUseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(applyRequest)
		err := uc.Apply(req.ApplicantEmailAddress)
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

func makeAddNextInterview(uc ScreeningUseCase) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(addNextInterviewRequest)
		err := uc.AddNextInterview(req.ScreeningID, req.InterviewDate)
		return addNextInterviewResponse{Err: err}, nil
	}
}
