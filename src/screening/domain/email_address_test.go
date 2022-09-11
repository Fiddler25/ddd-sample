package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmailAddress(t *testing.T) {
	type args struct {
		emailAddress string
	}
	tests := []struct {
		name      string
		args      args
		want      EmailAddress
		wantError bool
	}{
		{
			name:      "有効なメールアドレスの場合、インスタンスが生成されること",
			args:      args{emailAddress: "test@example.com"},
			want:      EmailAddress("test@example.com"),
			wantError: false,
		},
		{
			name:      "無効なメールアドレスの場合、空文字とエラーが返されること",
			args:      args{emailAddress: "test"},
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmailAddress(tt.args.emailAddress)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "引数が空文字の場合、trueを返すこと",
			args: args{value: ""},
			want: true,
		},
		{
			name: "引数が空文字ではない場合、falseを返すこと",
			args: args{value: "test"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isEmpty(tt.args.value)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsInvalidFormatEmailAddress(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "有効なメールアドレスの場合、falseを返すこと",
			args: args{email: "test@example.com"},
			want: false,
		},
		{
			name: "無効なメールアドレスの場合、trueを返すこと",
			args: args{email: "test"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInvalidFormatEmailAddress(tt.args.email)

			assert.Equal(t, tt.want, got)
		})
	}
}
