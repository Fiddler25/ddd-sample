package screening

import "github.com/google/uuid"

type ScreeningID struct {
	value string
}

func NewScreeningID() ScreeningID {
	return ScreeningID{value: uuid.NewString()}
}
