package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/ent/enttest"
	"ddd-sample/ent/schema/property"
	"ddd-sample/src/screening/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"testing"
)

func setup(t *testing.T) (*ent.Client, domain.ScreeningRepository, context.Context) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	return client, NewScreeningRepository(client), context.Background()
}

func TestScreeningRepository_Insert(t *testing.T) {
	client, repo, ctx := setup(t)
	defer client.Close()

	tests := []struct {
		name    string
		args    *domain.Screening
		want    *ent.Screening
		wantErr bool
	}{
		{
			name: "採用選考の登録が成功すること",
			args: &domain.Screening{
				ScreeningID:           domain.NewScreeningID(),
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
				Interviews:            domain.Interviews{},
			},
			want: &ent.Screening{
				ScreeningStatus:       property.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "ScreeningIDが50文字の場合、採用選考の登録が成功すること",
			args: &domain.Screening{
				ScreeningID:           domain.ScreeningID(strings.Repeat("a", 50)),
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
				Interviews:            domain.Interviews{},
			},
			want: &ent.Screening{
				ScreeningStatus:       property.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "ApplicantEmailAddressが50文字の場合、採用選考の登録が成功すること",
			args: &domain.Screening{
				ScreeningID:           domain.NewScreeningID(),
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: domain.EmailAddress(strings.Repeat("a", 38) + "@example.com"),
				Interviews:            domain.Interviews{},
			},
			want: &ent.Screening{
				ScreeningStatus:       property.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: strings.Repeat("a", 38) + "@example.com",
			},
			wantErr: false,
		},
		{
			name: "ScreeningIDが空の場合、エラーを返すこと",
			args: &domain.Screening{
				ScreeningID:           "",
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
				Interviews:            domain.Interviews{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ScreeningIDが51文字の場合、エラーを返すこと",
			args: &domain.Screening{
				ScreeningID:           domain.ScreeningID(strings.Repeat("a", 51)),
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
				Interviews:            domain.Interviews{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Statusが空の場合、エラーを返すこと",
			args: &domain.Screening{
				ScreeningID:           domain.NewScreeningID(),
				Status:                "",
				ApplyDate:             nil,
				ApplicantEmailAddress: "test@example.com",
				Interviews:            domain.Interviews{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ApplicantEmailAddressが空の場合、エラーを返すこと",
			args: &domain.Screening{
				ScreeningID:           domain.NewScreeningID(),
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: "",
				Interviews:            domain.Interviews{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ApplicantEmailAddressが51文字の場合、エラーを返すこと",
			args: &domain.Screening{
				ScreeningID:           domain.NewScreeningID(),
				Status:                domain.NotApplied,
				ApplyDate:             nil,
				ApplicantEmailAddress: domain.EmailAddress(strings.Repeat("a", 39) + "@example.com"),
				Interviews:            domain.Interviews{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Insert(ctx, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, _ := client.Screening.Get(ctx, string(tt.args.ScreeningID))

			opts := []cmp.Option{
				cmp.AllowUnexported(domain.Interviews{}),
				cmpopts.IgnoreFields(ent.Screening{}, "config", "ID", "Created", "Updated"),
			}
			if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
				t.Errorf("Insert() mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
