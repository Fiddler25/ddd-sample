package screening

import (
	"time"
)

type ID string

// Screening 採用選考
type Screening struct {
	ID                    ID         // 採用選考ID
	Status                Status     // 採用選考ステータス
	ApplyDate             *time.Time // 応募日
	ApplicantEmailAddress string     // 応募者メールアドレス
}

func newScreening(s *Screening) *Screening {
	return &Screening{
		ID:                    s.ID,
		Status:                s.Status,
		ApplyDate:             s.ApplyDate,
		ApplicantEmailAddress: s.ApplicantEmailAddress,
	}
}

type Status string

const (
	NotApplied                Status = "NOT_APPLIED"                 // 未応募
	DocumentScreening         Status = "DOCUMENT_SCREENING"          // 書類選考
	DocumentScreeningRejected Status = "DOCUMENT_SCREENING_REJECTED" // 書類不合格
	DocumentScreeningDeclined Status = "DOCUMENT_SCREENING_DECLINED" // 書類選考辞退
	InterviewScreening        Status = "INTERVIEW"                   // 面接選考中
	InterviewRejected         Status = "INTERVIEW_REJECTED"          // 面接不合格
	InterviewDeclined         Status = "INTERVIEW_DECLINED"          // 面接辞退
	Offered                   Status = "OFFERED"                     // 内定
	OfferDeclined             Status = "OFFER_DECLINED"              // 内定辞退
	Entered                   Status = "ENTERED"                     // 入社済
)
