package screening

import (
	"context"

	"github.com/google/uuid"
)

type StartFromPreInterviewOutput struct {
	ID                    ID
	Status                Status
	ApplicantEmailAddress string
}

// StartFromPreInterview 面談から新規候補者を登録する
func (s *service) StartFromPreInterview(ctx context.Context, applicantEmailAddress string) (*StartFromPreInterviewOutput, error) {
	e := &Screening{
		ID:                    ID(uuid.NewString()),
		Status:                NotApplied, // 面談からの場合はステータス「未応募」で登録
		ApplyDate:             nil,        // 未応募なので応募日はnil
		ApplicantEmailAddress: applicantEmailAddress,
	}
	model := newScreening(e)

	res, err := s.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	return &StartFromPreInterviewOutput{
		ID:                    res.ID,
		Status:                res.Status,
		ApplicantEmailAddress: res.ApplicantEmailAddress,
	}, nil
}
