package screening

import "time"

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
