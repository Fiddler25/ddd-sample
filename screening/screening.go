package screening

import "time"

type (
	ScreeningID     string
	ScreeningStatus string
	EmailAddress    string
)

// Screening 採用選考
type Screening struct {
	ScreeningID           ScreeningID     // 採用選考ID
	Status                ScreeningStatus // 採用選考ステータス
	ApplyDate             *time.Time      // 応募日
	ApplicantEmailAddress EmailAddress    // 応募者メールアドレス
}

func newScreening() *Screening {
	return &Screening{}
}
