package screening

import (
	"ddd-sample/src/domain/interview"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ID string
type status string

const (
	notApplied status = "NOT_APPLIED" // 未応募
	screening  status = "SCREENING"   // 面接選考中
	rejected   status = "REJECTED"    // 不合格
	passed     status = "PASSED"      // 合格
)

// Screening 採用選考
type Screening struct {
	screeningID           ID                    // 採用選考ID
	status                status                // 採用選考ステータス
	applyDate             *time.Time            // 応募日
	applicantEmailAddress string                // 応募者メールアドレス
	interviews            []interview.Interview // 面接
}

func newScreening() *Screening {
	return &Screening{}
}

// StartFromPreInterview 面談から採用選考を登録する際のファクトリメソッド
func StartFromPreInterview(applicantEmailAddress string) (*Screening, error) {
	if isEmpty(applicantEmailAddress) || isInvalidFormatEmailAddress(applicantEmailAddress) {
		return nil, fmt.Errorf("メールアドレスが正しくありません")
	}

	s := newScreening()

	s.screeningID = ID(uuid.NewString())
	s.status = notApplied // 面談からの場合はステータス「未応募」で登録
	s.applyDate = nil     // 未応募なので応募日はnull
	s.applicantEmailAddress = applicantEmailAddress
	s.interviews = []interview.Interview{}

	return s, nil
}

// Apply 面接から採用選考を登録する際のファクトリメソッド
func Apply(applicantEmailAddress string) (*Screening, error) {
	if isEmpty(applicantEmailAddress) || isInvalidFormatEmailAddress(applicantEmailAddress) {
		return nil, fmt.Errorf("メールアドレスが正しくありません")
	}

	s := newScreening()
	now := time.Now()

	s.screeningID = ID(uuid.NewString())
	s.status = screening // 面接からの場合はステータス「面接」で登録
	s.applyDate = &now   // 応募日は操作日付を使用
	s.applicantEmailAddress = applicantEmailAddress
	s.interviews = []interview.Interview{}

	return s, nil
}

// AddNextInterview 次の面接を設定する
func AddNextInterview(s *Screening, interviewDate time.Time) (*Screening, error) {
	if s.status != screening {
		return s, fmt.Errorf("不正な操作です")
	}

	nextInterviewNumber := len(s.interviews) + 1
	nextInterview := interview.NewInterview(interviewDate, nextInterviewNumber)
	s.interviews = append(s.interviews, nextInterview)

	return s, nil
}

func isEmpty(val string) bool {
	return val == ""
}

func isInvalidFormatEmailAddress(email string) bool {
	if email == "" {
		return true
	}

	// 何らかの処理
	return false
}
