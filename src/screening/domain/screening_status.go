package domain

import "fmt"

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

// NextStep 次のステップのステータスを取得する
func (s ScreeningStatus) NextStep() (ScreeningStatus, error) {
	switch s {
	case NotApplied:
		return DocumentScreening, nil
	case DocumentScreening:
		return InterviewScreening, nil
	case InterviewScreening:
		return Offered, nil
	case Offered:
		return Entered, nil
	default:
		return "", fmt.Errorf("許可されていない状態遷移です")
	}
}

// PreviousStep 「戻る」した時のステータスを取得する
func (s ScreeningStatus) PreviousStep() (ScreeningStatus, error) {
	switch s {
	case DocumentScreeningRejected:
		return DocumentScreening, nil
	case DocumentScreeningDeclined:
		return DocumentScreening, nil
	case InterviewScreening:
		return DocumentScreening, nil
	case InterviewRejected:
		return InterviewScreening, nil
	case InterviewDeclined:
		return InterviewScreening, nil
	case Offered:
		return InterviewScreening, nil
	case OfferDeclined:
		return Offered, nil
	case Entered:
		return Offered, nil
	default:
		return "", fmt.Errorf("許可されていない状態遷移です")
	}
}

// RejectStep 「不合格」した時のステータスを取得する
func (s ScreeningStatus) RejectStep() (ScreeningStatus, error) {
	switch s {
	case DocumentScreening:
		return DocumentScreeningRejected, nil
	case InterviewScreening:
		return InterviewRejected, nil
	default:
		return "", fmt.Errorf("許可されていない状態遷移です")
	}
}

// DeclineStep 「辞退」した時のステータスを取得する
func (s ScreeningStatus) DeclineStep() (ScreeningStatus, error) {
	switch s {
	case DocumentScreening:
		return DocumentScreeningDeclined, nil
	case InterviewScreening:
		return InterviewDeclined, nil
	case Offered:
		return OfferDeclined, nil
	default:
		return "", fmt.Errorf("許可されていない状態遷移です")
	}
}

func (s ScreeningStatus) CanAddInterview() bool {
	return s == InterviewScreening
}
