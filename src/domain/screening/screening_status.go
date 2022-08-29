package screening

type (
	ScreeningStatus struct {
		canAddInterview canAddInterview
	}

	canAddInterview bool
)

const (
	NotApplied                canAddInterview = false // 未応募
	DocumentScreening         canAddInterview = false // 書類選考
	DocumentScreeningRejected canAddInterview = false // 書類不合格
	DocumentScreeningDeclined canAddInterview = false // 書類選考辞退
	Interview                 canAddInterview = true  // 面接選考中
	InterviewRejected         canAddInterview = false // 面接不合格
	InterviewDeclined         canAddInterview = false // 面接辞退
	Offered                   canAddInterview = false // 内定
	OfferDeclined             canAddInterview = false // 内定辞退
	Entered                   canAddInterview = false // 入社済
)

func (s ScreeningStatus) CanAddInterview() bool {
	return bool(s.canAddInterview)
}
