package screening

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var now = time.Now()

const date = "2006-01-02"

func TestStartFromPreInterview(t *testing.T) {
	type args struct {
		applicantEmailAddress EmailAddress
	}
	tests := []struct {
		name string
		args args
		want *Screening
	}{
		{
			name: "採用選考を登録すると、ステータスが「未応募」・応募日がnullのインスタンスが生成されること",
			args: args{applicantEmailAddress: EmailAddress("test@example.com")},
			want: &Screening{
				Status:                NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: EmailAddress("test@example.com"),
				Interviews:            NewInterviews(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartFromPreInterview(tt.args.applicantEmailAddress)

			opts := []cmp.Option{
				cmp.AllowUnexported(Interviews{}),
				cmpopts.IgnoreFields(Screening{}, "ScreeningID"),
			}

			if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
				t.Errorf("Apply() mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		applicantEmailAddress EmailAddress
	}
	tests := []struct {
		name string
		args args
		want *Screening
	}{
		{
			name: "採用選考を登録すると、ステータスが「面接選考中」・応募日が本日のインスタンスが生成されること",
			args: args{applicantEmailAddress: EmailAddress("test@example.com")},
			want: &Screening{
				Status:                InterviewScreening,
				ApplicantEmailAddress: EmailAddress("test@example.com"),
				Interviews:            NewInterviews(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Apply(tt.args.applicantEmailAddress)

			opts := []cmp.Option{
				cmp.AllowUnexported(Interviews{}),
				cmpopts.IgnoreFields(Screening{}, "ScreeningID", "ApplyDate"),
			}

			if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
				t.Errorf("Apply() mismatch (-want, +got):\n%s", diff)
			}
			assert.Equal(t, now.Format(date), got.ApplyDate.Format(date))
		})
	}
}

func TestAddNextInterview(t *testing.T) {
	type args struct {
		interviewDate time.Time
	}
	tests := []struct {
		name      string
		args      args
		screening *Screening
		want      Interviews
		wantError bool
	}{
		{
			name: "有効な採用選考ステータスの場合、面接を追加すると面接次数がインクリメントされること",
			args: args{interviewDate: now},
			screening: &Screening{
				ScreeningID:           NewScreeningID(),
				Status:                InterviewScreening,
				ApplyDate:             nil,
				ApplicantEmailAddress: EmailAddress("test@example.com"),
				Interviews:            Interviews{},
			},
			want: Interviews{
				[]Interview{
					{
						interviewDate:       now,
						interviewNumber:     1,
						screeningStepResult: notEvaluated,
					},
				},
			},
			wantError: false,
		},
		{
			name: "無効な採用選考ステータスの場合、エラーを返すこと",
			args: args{interviewDate: now},
			screening: &Screening{
				ScreeningID:           NewScreeningID(),
				Status:                NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: EmailAddress("test@example.com"),
				Interviews:            Interviews{},
			},
			want:      Interviews{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.screening.AddNextInterview(tt.args.interviewDate)

			want := tt.want.interviews
			got := tt.screening.Interviews.interviews
			opts := []cmp.Option{
				cmp.AllowUnexported(Interview{}),
				cmpopts.IgnoreFields(Interview{}, "interviewID"),
			}

			if diff := cmp.Diff(want, got, opts...); diff != "" {
				t.Errorf("Interviews mismatch (-want, +got):\n%s", diff)
			}
			if err != nil {
				assert.EqualError(t, err, "不正な操作です")
			}
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestStepToNext(t *testing.T) {
	tests := []struct {
		name      string
		screening *Screening
		want      ScreeningStatus
		wantError bool
	}{
		{
			name: "許可されたステータスの場合、次のステップのステータスが設定されること",
			screening: &Screening{
				ScreeningID:           NewScreeningID(),
				Status:                NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: EmailAddress("test@example.com"),
				Interviews:            Interviews{},
			},
			want:      DocumentScreening,
			wantError: false,
		},
		{
			name: "許可されていないステータスの場合、エラーを返すこと",
			screening: &Screening{
				ScreeningID:           NewScreeningID(),
				Status:                InterviewRejected,
				ApplyDate:             nil,
				ApplicantEmailAddress: EmailAddress("test@example.com"),
				Interviews:            Interviews{},
			},
			want:      InterviewRejected,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.screening.StepToNext()

			assert.Equal(t, tt.want, tt.screening.Status)
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}
