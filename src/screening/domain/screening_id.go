package domain

import (
	"net/url"

	"github.com/google/uuid"
)

type ScreeningID string

func NewScreeningID() ScreeningID {
	url.JoinPath("https://sample.com", "api")

	return ScreeningID(uuid.NewString())
}
