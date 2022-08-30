package screening

import (
	"ddd-sample/src/domain/screening"
	"time"
)

type ScreeningUseCase struct {
	screening screening.ScreeningRepository
}

func NewScreeningUseCase(screening screening.ScreeningRepository) ScreeningUseCase {
	return ScreeningUseCase{screening: screening}
}

// StartFromPreInterview 面談から新規候補者を登録する
func (uc ScreeningUseCase) StartFromPreInterview(applicantEmailAddress string) error {
	e, err := screening.NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s, err := screening.StartFromPreInterview(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(s)
}

// Apply 新規応募者を登録する
func (uc ScreeningUseCase) Apply(applicantEmailAddress string) error {
	e, err := screening.NewEmailAddress(applicantEmailAddress)
	if err != nil {
		return err
	}

	s, err := screening.Apply(e)
	if err != nil {
		return err
	}

	return uc.screening.Insert(s)
}

// AddNextInterview 次の面接を設定する
func (uc ScreeningUseCase) AddNextInterview(screeningID string, interviewDate time.Time) error {
	s, err := uc.screening.FindByID(screening.ScreeningID(screeningID))
	if err != nil {
		return err
	}

	if err := s.AddNextInterview(interviewDate); err != nil {
		return err
	}

	return uc.screening.Update(s)
}

// StepToNext 採用選考を次のステップに進める
func (uc ScreeningUseCase) StepToNext(screeningID string) error {
	s, err := uc.screening.FindByID(screening.ScreeningID(screeningID))
	if err != nil {
		return err
	}

	s.StepToNext()
	return uc.screening.Update(s)
}
