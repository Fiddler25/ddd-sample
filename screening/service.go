package screening

import "context"

type service struct {
	repo Repository
}

type Service interface {
	StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
