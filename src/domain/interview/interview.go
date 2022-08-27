package interview

import (
	"ddd-sample/src/domain/screening"
	"time"
)

type ID string
type ScreeningStepResult string

const (
	NotEvaluated ScreeningStepResult = "NOT_EVALUATED" // 未評価
	Pass         ScreeningStepResult = "PASS"          // 合格
	Fail         ScreeningStepResult = "FAIL"          // 不合格
)

type Interview struct {
	InterviewID         ID                  // 面接ID
	ScreeningID         screening.ID        // 採用選考ID
	ScreeningDate       time.Time           // 選考日
	InterviewNumber     int                 // 面接次数
	ScreeningStepResult ScreeningStepResult // 面接結果
	RecruiterID         int                 // 採用担当者ID
}
