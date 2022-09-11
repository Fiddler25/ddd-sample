package screening

import (
	"fmt"
	"net/mail"
)

type EmailAddress string

func NewEmailAddress(value string) (EmailAddress, error) {
	if isEmpty(value) || isInvalidFormatEmailAddress(value) {
		return "", fmt.Errorf("メールアドレスが正しくありません")
	}

	return EmailAddress(value), nil
}

func isEmpty(value string) bool {
	return value == ""
}

func isInvalidFormatEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}
