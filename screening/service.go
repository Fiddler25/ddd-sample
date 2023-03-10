package screening

import "context"

type service struct {
	repo Repository
}

type Service interface {
	StartFromPreInterview(ctx context.Context, applicantEmailAddress string) (*StartFromPreInterviewOutput, error)
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

var _ Service = (*MockService)(nil)

type MockService struct {
	StartFromPreInterviewFunc func(ctx context.Context, applicantEmailAddress string) (*StartFromPreInterviewOutput, error)
}

func (m *MockService) StartFromPreInterview(ctx context.Context, applicantEmailAddress string) (*StartFromPreInterviewOutput, error) {
	return m.StartFromPreInterviewFunc(ctx, applicantEmailAddress)
}
