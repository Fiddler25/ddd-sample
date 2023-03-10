package screening_test

import (
	"context"
	"ddd-sample/ent/enttest"
	"ddd-sample/screening"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
)

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	repo := screening.NewRepository(client)

	type (
		give struct {
			screening *screening.Screening
		}

		want struct {
			screening *screening.Screening
			err       bool
		}
	)
	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "【OK】登録成功",
			give: give{screening: &screening.Screening{
				ID:                    screening.ID("1"),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
			want: want{screening: &screening.Screening{
				ID:                    screening.ID("1"),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
		},
		{
			name: "【OK】IDが50文字",
			give: give{screening: &screening.Screening{
				ID:                    screening.ID(strings.Repeat("a", 50)),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
			want: want{screening: &screening.Screening{
				ID:                    screening.ID(strings.Repeat("a", 50)),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
		},
		{
			name: "【OK】ApplicantEmailAddressが50文字",
			give: give{screening: &screening.Screening{
				ID:                    screening.ID("email-50"),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: strings.Repeat("a", 50),
			}},
			want: want{screening: &screening.Screening{
				ID:                    screening.ID("email-50"),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: strings.Repeat("a", 50),
			}},
		},
		{
			name: "【NG】IDが51文字",
			give: give{screening: &screening.Screening{
				ID:                    screening.ID(strings.Repeat("a", 51)),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: "test@example.com",
			}},
			want: want{err: true},
		},
		{
			name: "【NG】ApplicantEmailAddressが51文字",
			give: give{screening: &screening.Screening{
				ID:                    screening.ID("email-51"),
				Status:                screening.NotApplied,
				ApplicantEmailAddress: strings.Repeat("a", 51),
			}},
			want: want{err: true},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Create(context.Background(), tt.give.screening)

			if diff := cmp.Diff(tt.want.screening, got); diff != "" {
				t.Errorf("screening mismatch (-want, +got):\n%s", diff)
			}

			if (err != nil) != tt.want.err {
				t.Errorf("unexpected error = %v", err)
			}
		})
	}
}
