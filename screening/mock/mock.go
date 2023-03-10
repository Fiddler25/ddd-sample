package mock

import (
	"context"
	"ddd-sample/screening"
)

var _ screening.Service = (*ScreeningService)(nil)

type ScreeningService struct {
	StartFromPreInterviewFunc func(ctx context.Context, applicantEmailAddress string) (*screening.StartFromPreInterviewOutput, error)
}

func (mock *ScreeningService) StartFromPreInterview(ctx context.Context, applicantEmailAddress string) (*screening.StartFromPreInterviewOutput, error) {
	return mock.StartFromPreInterviewFunc(ctx, applicantEmailAddress)
}
