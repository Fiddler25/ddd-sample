package screening

import (
	"ddd-sample/src/domain/interview"
	"ddd-sample/src/domain/screening"
	"fmt"
	"github.com/google/uuid"
	"time"
)

// 面談から新規候補者を登録する
func startFromPreInterview(applicantEmailAddress string) error {
	if isEmpty(applicantEmailAddress) || isInvalidFormatEmailAddress(applicantEmailAddress) {
		return fmt.Errorf("メールアドレスが正しくありません")
	}

	s := new(screening.Screening)

	s.ScreeningID = screening.ID(uuid.NewString())
	s.Status = screening.NotApplied // 面談からの場合はステータス「未応募」で登録
	s.ApplyDate = nil               // 未応募なので応募日はnull
	s.ApplicantEmailAddress = applicantEmailAddress

	// insert
	return nil
}

// 新規応募者を登録する
func apply(applicantEmailAddress string) error {
	if isEmpty(applicantEmailAddress) || isInvalidFormatEmailAddress(applicantEmailAddress) {
		return fmt.Errorf("メールアドレスが正しくありません")
	}

	s := &screening.Screening{}
	now := time.Now()

	s.ScreeningID = screening.ID(uuid.NewString())
	s.Status = screening.Interview // 面接からの場合はステータス「面接」で登録
	s.ApplyDate = &now             // 応募日は操作日付を使用
	s.ApplicantEmailAddress = applicantEmailAddress

	// insert
	return nil
}

func isEmpty(val string) bool {
	return val == ""
}

func isInvalidFormatEmailAddress(email string) bool {
	if email == "" {
		return false
	}

	// .....
	return true
}

// 次の面接を設定する
func addNextInterview(screeningID screening.ID, interviewDate time.Time) error {
	// 保存されている採用選考オブジェクトを取得したと仮定
	s := &screening.Screening{}

	if s.Status != screening.Interview {
		return fmt.Errorf("不正な操作です")
	}

	// 保存されている面接オブジェクトの一覧を取得したと仮定
	var is []interview.Interview
	i := &interview.Interview{}

	i.InterviewID = interview.ID(uuid.NewString())
	i.ScreeningID = screeningID
	i.InterviewNumber = len(is) + 1 // 面接次数は保存されているインタビューの数+1とする
	i.ScreeningDate = interviewDate

	// insert
	return nil
}
