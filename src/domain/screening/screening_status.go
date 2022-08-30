package screening

type ScreeningStatus string

const (
	NotApplied                ScreeningStatus = "NOT_APPLIED"                 // 未応募
	DocumentScreening         ScreeningStatus = "DOCUMENT_SCREENING"          // 書類選考
	DocumentScreeningRejected ScreeningStatus = "DOCUMENT_SCREENING_REJECTED" // 書類不合格
	DocumentScreeningDeclined ScreeningStatus = "DOCUMENT_SCREENING_DECLINED" // 書類選考辞退
	InterviewScreening        ScreeningStatus = "INTERVIEW"                   // 面接選考中
	InterviewRejected         ScreeningStatus = "INTERVIEW_REJECTED"          // 面接不合格
	InterviewDeclined         ScreeningStatus = "INTERVIEW_DECLINED"          // 面接辞退
	Offered                   ScreeningStatus = "OFFERED"                     // 内定
	OfferDeclined             ScreeningStatus = "OFFER_DECLINED"              // 内定辞退
	Entered                   ScreeningStatus = "ENTERED"                     // 入社済
)

func (s ScreeningStatus) CanAddInterview() bool {
	return s.canAddInterview()
}

func (s ScreeningStatus) canAddInterview() bool {
	if s == InterviewScreening {
		return true
	}
	return false
}
