package vo

import "fmt"

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
	// 何らかの処理
	return false
}
