package screening

import "time"

type ID string
type Status string

const (
	NotApplied Status = "NOT_APPLIED" // 未応募
	Interview  Status = "INTERVIEW"   // 面接選考中
	Refected   Status = "REFECTED"    // 不合格
	Passed     Status = "PASSED"      // 合格
)

// Screening 採用選考
type Screening struct {
	ScreeningID           ID         // 採用選考ID
	Status                Status     // 採用選考ステータス
	ApplyDate             *time.Time // 応募日
	ApplicantEmailAddress string     // 応募者メールアドレス
}
