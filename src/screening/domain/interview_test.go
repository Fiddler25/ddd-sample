package domain

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

			opts := []cmp.Option{
				cmp.AllowUnexported(Interview{}),
				cmpopts.IgnoreFields(Interview{}, "interviewID"),
			}

			if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
				t.Errorf("NewInterview() mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
