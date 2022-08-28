package interview

import (
	"ddd-sample/src/domain/screening"
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
