package screening

import "github.com/google/uuid"

type ScreeningID string

func NewScreeningID() ScreeningID {
	return ScreeningID(uuid.NewString())
}
