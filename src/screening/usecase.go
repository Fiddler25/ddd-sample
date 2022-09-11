package screening

import (
	"time"
)

type Usecase interface {
	StartFromPreInterview(applicantEmailAddress string) error
	Apply(applicantEmailAddress string) error
	AddNextInterview(screeningID string, interviewDate time.Time) error
	StepToNext(screeningID string) error
}

type usecase struct {
	screening ScreeningRepository
}

func NewUsecase(screening ScreeningRepository) Usecase {
	return &usecase{screening: screening}
}

// StartFromPreInterview 面談から新規候補者を登録する
func (uc usecase) StartFromPreInterview(applicantEmailAddress string) error {
	e, err := NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s, err := StartFromPreInterview(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(s)
}

// Apply 新規応募者を登録する
func (uc usecase) Apply(applicantEmailAddress string) error {
	e, err := NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s, err := Apply(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(s)
}

// AddNextInterview 次の面接を設定する
func (uc usecase) AddNextInterview(screeningID string, interviewDate time.Time) error {
	s, err := uc.screening.FindByID(ScreeningID(screeningID))
	if err != nil {
		return err
	}

	if err := s.AddNextInterview(interviewDate); err != nil {
		return err
	}

	return uc.screening.Update(s)
}

// StepToNext 採用選考を次のステップに進める
func (uc usecase) StepToNext(screeningID string) error {
	s, err := uc.screening.FindByID(ScreeningID(screeningID))
	if err != nil {
		return err
	}

	s.StepToNext()
	return uc.screening.Update(s)
}
