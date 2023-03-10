package screening_test

import (
	"context"
	"ddd-sample/screening"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMakeStartFromPreInterviewEndpoint(t *testing.T) {
	t.Parallel()

	mock := &screening.MockService{}

	type give struct {
		req                   screening.StartFromPreInterviewRequest
		id                    screening.ID
		status                screening.Status
		applicantEmailAddress string
		err                   error
	}

	tests := []struct {
		name string
		give give
		want screening.StartFromPreInterviewResponse
	}{
		{
			name: "【OK】正常リクエストが成功する",
			give: give{
				req:                   screening.StartFromPreInterviewRequest{ApplicantEmailAddress: "test@example.com"},
				id:                    "screening-id",
				status:                screening.NotApplied,
				applicantEmailAddress: "test@example.com",
			},
			want: screening.StartFromPreInterviewResponse{
				ID:                    "screening-id",
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			},
		},
		{
			name: "【NG】メールアドレスが空文字",
			give: give{req: screening.StartFromPreInterviewRequest{ApplicantEmailAddress: ""}},
			want: screening.StartFromPreInterviewResponse{Err: screening.ErrValidation},
		},
		{
			name: "【NG】メールアドレスの形式が不正",
			give: give{req: screening.StartFromPreInterviewRequest{ApplicantEmailAddress: "testexample.com"}},
			want: screening.StartFromPreInterviewResponse{Err: screening.ErrValidation},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mock.StartFromPreInterviewFunc = func(ctx context.Context, applicantEmailAddress string) (*screening.StartFromPreInterviewOutput, error) {
				return &screening.StartFromPreInterviewOutput{
					ID:                    tt.give.id,
					Status:                tt.give.status,
					ApplicantEmailAddress: tt.give.applicantEmailAddress,
				}, tt.give.err
			}

			resp, _ := screening.MakeStartFromPreInterviewEndpoint(mock)(context.Background(), tt.give.req)
			got, ok := resp.(screening.StartFromPreInterviewResponse)
			if !ok {
				t.Errorf("unexpected response = %v", resp)
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(screening.StartFromPreInterviewResponse{}, "Err")); diff != "" {
				t.Errorf("response mismatch (-want +got)\n%s", diff)
			}

			if !errors.Is(got.Err, tt.want.Err) {
				t.Errorf("err = %v, want = %v", got.Err, tt.want.Err)
			}
		})
	}
}
