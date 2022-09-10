package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/ent/schema/property"
)

type screeningRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewScreeningRepository(ctx context.Context, client *ent.Client) ScreeningRepository {
	return &screeningRepository{
		ctx:    ctx,
		client: client,
	}
}

func (r screeningRepository) FindByID(screeningID ScreeningID) (*Screening, error) {
	s, err := r.client.Screening.Get(r.ctx, string(screeningID))
	if err != nil {
		return nil, err
	}
	return reconstruct(s), nil
}

func reconstruct(s *ent.Screening) *Screening {
	return &Screening{
		ScreeningID:           ScreeningID(s.ID),
		Status:                ScreeningStatus(s.ScreeningStatus),
		ApplyDate:             s.ApplyDate,
		ApplicantEmailAddress: EmailAddress(s.ApplicantEmailAddress),
	}
}

func (r screeningRepository) Insert(s *Screening) error {
	if _, err := r.client.Screening.
		Create().
		SetID(string(s.ScreeningID)).
		SetScreeningStatus(property.ScreeningStatus(s.Status)).
		SetNillableApplyDate(s.ApplyDate).
		SetApplicantEmailAddress(string(s.ApplicantEmailAddress)).
		Save(r.ctx); err != nil {
		return err
	}
	return nil
}

func (r screeningRepository) Update(screening *Screening) error {
	// 更新処理
	return nil
}
