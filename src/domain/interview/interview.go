package interview

import (
	"github.com/google/uuid"
	"time"
)

type ID string
type screeningStepResult string

const (
	notEvaluated screeningStepResult = "NOT_EVALUATED" // 未評価
	pass         screeningStepResult = "PASS"          // 合格
	fail         screeningStepResult = "FAIL"          // 不合格
)

// Interview 面接
type Interview struct {
	interviewID         ID                  // 面接ID
	interviewDate       time.Time           // 選考日
	interviewNumber     int                 // 面接次数
	screeningStepResult screeningStepResult // 面接結果
}

func NewInterview(interviewDate time.Time, interviewNumber int) Interview {
	return Interview{
		interviewID:         ID(uuid.NewString()),
		interviewDate:       interviewDate,
		interviewNumber:     interviewNumber,
		screeningStepResult: notEvaluated,
	}
}
