package screening

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextStep(t *testing.T) {
	tests := []struct {
		name      string
		status    ScreeningStatus
		want      ScreeningStatus
		wantError bool
	}{
		{
			name:      "ステータスが「未応募」の場合、「書類選考」を返すこと",
			status:    NotApplied,
			want:      DocumentScreening,
			wantError: false,
		},
		{
			name:      "ステータスが「書類選考」の場合、「面接選考中」を返すこと",
			status:    DocumentScreening,
			want:      InterviewScreening,
			wantError: false,
		},
		{
			name:      "ステータスが「面接選考中」の場合、「内定」を返すこと",
			status:    InterviewScreening,
			want:      Offered,
			wantError: false,
		},
		{
			name:      "ステータスが「内定」の場合、「入社済」を返すこと",
			status:    Offered,
			want:      Entered,
			wantError: false,
		},
		{
			name:      "ステータスが「書類不合格」の場合、空文字とエラーを返すこと",
			status:    DocumentScreeningRejected,
			want:      "",
			wantError: true,
		},
		{
			name:      "ステータスが「書類選考辞退」の場合、空文字とエラーを返すこと",
			status:    DocumentScreeningDeclined,
			want:      "",
			wantError: true,
		},
		{
			name:      "ステータスが「面接不合格」の場合、空文字とエラーを返すこと",
			status:    InterviewRejected,
			want:      "",
			wantError: true,
		},
		{
			name:      "ステータスが「面接辞退」の場合、空文字とエラーを返すこと",
			status:    InterviewDeclined,
			want:      "",
			wantError: true,
		},
		{
			name:      "ステータスが「内定辞退」の場合、空文字とエラーを返すこと",
			status:    OfferDeclined,
			want:      "",
			wantError: true,
		},
		{
			name:      "ステータスが「入社済」の場合、空文字とエラーを返すこと",
			status:    Entered,
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := tt.status.NextStep()

			assert := assert.New(t)
			assert.Equal(tt.want, s)
			assert.Equal(tt.wantError, err != nil)
			if err != nil {
				assert.EqualError(err, "許可されていない状態遷移です")
			}
		})
	}
}
