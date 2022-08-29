package screening

import (
	"ddd-sample/src/domain/interview"
	"ddd-sample/src/domain/vo"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ID string

// Screening 採用選考
type Screening struct {
	screeningID           ID                    // 採用選考ID
	status                ScreeningStatus       // 採用選考ステータス
	applyDate             *time.Time            // 応募日
	applicantEmailAddress vo.EmailAddress       // 応募者メールアドレス
	interviews            []interview.Interview // 面接
}

func newScreening() *Screening {
	return &Screening{}
}

// StartFromPreInterview 面談から採用選考を登録する際のファクトリメソッド
func StartFromPreInterview(applicantEmailAddress vo.EmailAddress) (*Screening, error) {
	s := newScreening()

	s.screeningID = ID(uuid.NewString())
	s.status = ScreeningStatus{NotApplied} // 面談からの場合はステータス「未応募」で登録
	s.applyDate = nil                      // 未応募なので応募日はnull
	s.applicantEmailAddress = applicantEmailAddress
	s.interviews = []interview.Interview{}

	return s, nil
}

// Apply 面接から採用選考を登録する際のファクトリメソッド
func Apply(applicantEmailAddress vo.EmailAddress) (*Screening, error) {
	s := newScreening()
	now := time.Now()

	s.screeningID = ID(uuid.NewString())
	s.status = ScreeningStatus{Interview} // 面接からの場合はステータス「面接」で登録
	s.applyDate = &now                    // 応募日は操作日付を使用
	s.applicantEmailAddress = applicantEmailAddress
	s.interviews = []interview.Interview{}

	return s, nil
}

// AddNextInterview 次の面接を設定する
func (s *Screening) AddNextInterview(interviewDate time.Time) (*Screening, error) {
	if !s.status.CanAddInterview() {
		return s, fmt.Errorf("不正な操作です")
	}

	nextInterviewNumber := len(s.interviews) + 1
	nextInterview := interview.NewInterview(interviewDate, nextInterviewNumber)
	s.interviews = append(s.interviews, nextInterview)

	return s, nil
}
