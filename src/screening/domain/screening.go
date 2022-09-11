package domain

import (
	"fmt"
	"time"
)

// Screening 採用選考
type Screening struct {
	ScreeningID           ScreeningID     // 採用選考ID
	Status                ScreeningStatus // 採用選考ステータス
	ApplyDate             *time.Time      // 応募日
	ApplicantEmailAddress EmailAddress    // 応募者メールアドレス
	Interviews            Interviews      // 面接
}

func newScreening() *Screening {
	return &Screening{}
}

// StartFromPreInterview 面談から採用選考を登録する際のファクトリメソッド
func StartFromPreInterview(applicantEmailAddress EmailAddress) *Screening {
	s := newScreening()

	s.ScreeningID = NewScreeningID()
	s.Status = NotApplied // 面談からの場合はステータス「未応募」で登録
	s.ApplyDate = nil     // 未応募なので応募日はnull
	s.ApplicantEmailAddress = applicantEmailAddress
	s.Interviews = NewInterviews()

	return s
}

// Apply 面接から採用選考を登録する際のファクトリメソッド
func Apply(applicantEmailAddress EmailAddress) *Screening {
	s := newScreening()
	now := time.Now()

	s.ScreeningID = NewScreeningID()
	s.Status = InterviewScreening // 面接からの場合はステータス「面接選考中」で登録
	s.ApplyDate = &now            // 応募日は登録日を使用
	s.ApplicantEmailAddress = applicantEmailAddress
	s.Interviews = NewInterviews()

	return s
}

// AddNextInterview 次の面接を設定する
func (s *Screening) AddNextInterview(interviewDate time.Time) error {
	if !s.Status.CanAddInterview() {
		return fmt.Errorf("不正な操作です")
	}

	s.Interviews.AddNextInterview(interviewDate)

	return nil
}

func (s *Screening) StepToNext() error {
	status, err := s.Status.NextStep()
	if err != nil {
		return err
	}
	s.Status = status
	return nil
}
