package screening

import (
	"ddd-sample/src/domain/screening"
	"ddd-sample/src/domain/vo"
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
	e, err := vo.NewEmailAddress(applicantEmailAddress)
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
	e, err := vo.NewEmailAddress(applicantEmailAddress)
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
func (uc ScreeningUseCase) AddNextInterview(screeningID screening.ID, interviewDate time.Time) error {
	s, err := uc.screening.FindByID(screeningID)
	if err != nil {
		return err
	}
	screening.AddNextInterview(s, interviewDate)

	return uc.screening.Update(s)
}
