package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/ent/schema/property"
	"ddd-sample/src/screening/domain"
)

type screeningRepository struct {
	client *ent.Client
}

func NewScreeningRepository(client *ent.Client) domain.ScreeningRepository {
	return &screeningRepository{client: client}
}

func (r screeningRepository) FindByID(ctx context.Context, screeningID domain.ScreeningID) (*domain.Screening, error) {
	s, err := r.client.Screening.Get(ctx, string(screeningID))
	if err != nil {
		return nil, err
	}
	return reconstruct(s), nil
}

func reconstruct(s *ent.Screening) *domain.Screening {
	return &domain.Screening{
		ScreeningID:           domain.ScreeningID(s.ID),
		Status:                domain.ScreeningStatus(s.ScreeningStatus),
		ApplyDate:             s.ApplyDate,
		ApplicantEmailAddress: domain.EmailAddress(s.ApplicantEmailAddress),
	}
}

func (r screeningRepository) Insert(ctx context.Context, s *domain.Screening) error {
	if _, err := r.client.Screening.
		Create().
		SetID(string(s.ScreeningID)).
		SetScreeningStatus(property.ScreeningStatus(s.Status)).
		SetNillableApplyDate(s.ApplyDate).
		SetApplicantEmailAddress(string(s.ApplicantEmailAddress)).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (r screeningRepository) Update(ctx context.Context, screening *domain.Screening) error {
	// 更新処理
	return nil
}
