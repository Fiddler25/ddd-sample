package screening

import (
	"context"
	"ddd-sample/src/screening/domain"
	"time"
)

type Usecase interface {
	StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error
	Apply(ctx context.Context, applicantEmailAddress string) error
	AddNextInterview(ctx context.Context, screeningID string, interviewDate time.Time) error
	StepToNext(ctx context.Context, screeningID string) error
}

type usecase struct {
	screening domain.ScreeningRepository
}

func NewUsecase(screening domain.ScreeningRepository) Usecase {
	return &usecase{screening: screening}
}

// StartFromPreInterview 面談から新規候補者を登録する
func (uc usecase) StartFromPreInterview(ctx context.Context, applicantEmailAddress string) error {
	e, err := domain.NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s := domain.StartFromPreInterview(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(ctx, s)
}

// Apply 新規応募者を登録する
func (uc usecase) Apply(ctx context.Context, applicantEmailAddress string) error {
	e, err := domain.NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s := domain.Apply(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(ctx, s)
}

// AddNextInterview 次の面接を設定する
func (uc usecase) AddNextInterview(ctx context.Context, screeningID string, interviewDate time.Time) error {
	s, err := uc.screening.FindByID(ctx, domain.ScreeningID(screeningID))
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
	s, err := uc.screening.FindByID(ctx, domain.ScreeningID(screeningID))
	if err != nil {
		return err
	}

	s.StepToNext()
	return uc.screening.Update(ctx, s)
}
