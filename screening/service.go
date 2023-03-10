package screening

import "context"

type service struct {
}

type Service interface {
	StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error
}

func NewService() Service {
	return &service{}
}
