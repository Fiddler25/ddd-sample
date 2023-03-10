package screening_test

import (
	"context"
	"ddd-sample/screening"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestService_StartFromPreInterview(t *testing.T) {
	t.Parallel()

	mock := &screening.MockRepository{}

	type (
		give struct {
			screening *screening.Screening
			err       error
		}

		want struct {
			screening *screening.StartFromPreInterviewOutput
			err       bool
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "【OK】正常終了",
			give: give{screening: &screening.Screening{
				ID:                    "1",
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
			want: want{screening: &screening.StartFromPreInterviewOutput{
				ID:                    "1",
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
		},
		{
			name: "【NG】StartFromPreInterviewでエラー発生",
			give: give{screening: &screening.Screening{}, err: errors.New("dummy-error")},
			want: want{err: true},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mock.CreateFunc = func(ctx context.Context, s *screening.Screening) (*screening.Screening, error) {
				return &screening.Screening{
					ID:                    tt.give.screening.ID,
					Status:                tt.give.screening.Status,
					ApplicantEmailAddress: tt.give.screening.ApplicantEmailAddress,
				}, tt.give.err
			}
			s := screening.NewService(mock)

			got, err := s.StartFromPreInterview(context.Background(), "test@example.com")

			if diff := cmp.Diff(tt.want.screening, got); diff != "" {
				t.Errorf("screening mismatch (-want, +got):\n%s", diff)
			}

			if (err != nil) != tt.want.err {
				t.Errorf("unexpected error = %v", err)
			}
		})
	}
}
