package screening

import (
	"context"

	"github.com/google/uuid"
)

func (s *service) StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error {
	e := &Screening{
		ID:                    ID(uuid.NewString()),
		Status:                NotApplied,
		ApplicantEmailAddress: applicantEmailAddress,
	}
	model := newScreening(e)

	return s.repo.Insert(ctx, model)

}
