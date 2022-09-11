package screening

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUsecase_StartFromPreInterview(t *testing.T) {
	type args struct {
		applicantEmailAddress string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "処理が成功すること",
			args:    args{applicantEmailAddress: "test@example.com"},
			wantErr: false,
			err:     nil,
		},
		{
			name:    "処理が失敗すること",
			args:    args{applicantEmailAddress: ""},
			wantErr: true,
			err:     errors.New(""),
		},
	}

	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		uc := NewMockUsecase(ctrl)

		uc.EXPECT().StartFromPreInterview(gomock.Any(), tt.args.applicantEmailAddress).Return(tt.err)
		err := uc.StartFromPreInterview(context.TODO(), tt.args.applicantEmailAddress)

		if (err != nil) != tt.wantErr {
			t.Errorf("StartFromPreInterview() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
