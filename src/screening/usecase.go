package screening

import (
	"context"
	"time"
)

type Usecase interface {
	StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error
	Apply(ctx context.Context, applicantEmailAddress string) error
	AddNextInterview(ctx context.Context, screeningID string, interviewDate time.Time) error
	StepToNext(ctx context.Context, screeningID string) error
}

type usecase struct {
	screening ScreeningRepository
}

func NewUsecase(screening ScreeningRepository) Usecase {
	return &usecase{screening: screening}
}

// StartFromPreInterview 面談から新規候補者を登録する
func (uc usecase) StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error {
	e, err := NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s := StartFromPreInterview(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(ctx, s)
}

// Apply 新規応募者を登録する
func (uc usecase) Apply(ctx context.Context, applicantEmailAddress string) error {
	e, err := NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s := Apply(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(ctx, s)
}

// AddNextInterview 次の面接を設定する
func (uc usecase) AddNextInterview(ctx context.Context, screeningID string, interviewDate time.Time) error {
	s, err := uc.screening.FindByID(ctx, ScreeningID(screeningID))
	if err != nil {
		return err
	}

	if err := s.AddNextInterview(interviewDate); err != nil {
		return err
	}

	return uc.screening.Update(ctx, s)
}

// StepToNext 採用選考を次のステップに進める
func (uc usecase) StepToNext(ctx context.Context, screeningID string) error {
	s, err := uc.screening.FindByID(ctx, ScreeningID(screeningID))
	if err != nil {
		return err
	}

	s.StepToNext()
	return uc.screening.Update(ctx, s)
}
