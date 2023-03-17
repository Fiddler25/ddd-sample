package domain

import (
	"log"
	"net/url"

	"github.com/google/uuid"
)

type ScreeningID string

func NewScreeningID() ScreeningID {
	_, err := url.JoinPath("https://sample.com", "api")
	if err != nil {
		log.Println(err)
	}

	return ScreeningID(uuid.NewString())
}
