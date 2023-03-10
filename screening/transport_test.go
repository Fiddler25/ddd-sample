package screening_test

import (
	"bytes"
	"context"
	"ddd-sample/screening"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestStartFromPreInterviewHandler(t *testing.T) {
	t.Parallel()

	var (
		mock = &screening.MockService{}
		id   = screening.ID(uuid.NewString())
	)

	type (
		give struct {
			id                    screening.ID
			Status                screening.Status
			ApplicantEmailAddress string
			err                   error
		}

		want struct {
			ID                    screening.ID
			Status                screening.Status
			ApplicantEmailAddress string
			statusCode            int
			err                   error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "【OK】正常リクエストが成功する",
			give: give{
				id:                    id,
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			},
			want: want{
				ID:                    id,
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
				statusCode:            http.StatusOK,
			},
		},
		{
			name: "【NG】StartFromPreInterview()でエラー発生",
			give: give{err: errors.New("dummy-error")},
			want: want{statusCode: http.StatusInternalServerError, err: errors.New("dummy-error")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			body := &screening.StartFromPreInterviewRequest{ApplicantEmailAddress: "test@example.com"}
			b, err := json.Marshal(body)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest("POST", "/screening/v1/start-from-pre-interview", bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}

			mock.StartFromPreInterviewFunc = func(ctx context.Context, applicantEmailAddress string) (*screening.StartFromPreInterviewOutput, error) {
				return &screening.StartFromPreInterviewOutput{
					ID:                    tt.give.id,
					Status:                tt.give.Status,
					ApplicantEmailAddress: tt.give.ApplicantEmailAddress,
				}, tt.give.err
			}

			r := screening.MakeHandler(mock)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			want := &screening.StartFromPreInterviewResponse{
				ID:                    tt.want.ID,
				Status:                tt.want.Status,
				ApplicantEmailAddress: tt.want.ApplicantEmailAddress,
			}

			var got *screening.StartFromPreInterviewResponse
			if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("response mismatch (-want +got)\n%s", diff)
			}

			if resp.StatusCode != tt.want.statusCode {
				t.Errorf("status code want = %d, got = %d", tt.want.statusCode, resp.StatusCode)
			}
		})
	}
}
