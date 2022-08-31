package screening

import (
	"testing"
	"time"
)

func TestNewInterview(t *testing.T) {
	type args struct {
		interviewDate   time.Time
		interviewNumber int
	}
	now := time.Now()

	tests := []struct {
		name string
		args args
		want Interview
	}{
		{
			name: "新しく面接を作成すると、未評価のインスタンスが生成されること",
			args: args{
				interviewDate:   now,
				interviewNumber: 0,
			},
			want: Interview{
				interviewDate:       now,
				interviewNumber:     0,
				screeningStepResult: notEvaluated,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInterview(tt.args.interviewDate, tt.args.interviewNumber)
			tt.want.interviewID = got.interviewID

			if got != tt.want {
				t.Errorf("NewInterview() = %v, want %v", got, tt.want)
			}
		})
	}
}
