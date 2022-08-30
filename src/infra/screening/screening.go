package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/ent/schema/property"
	"ddd-sample/src/domain/screening"
)

type screeningRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewScreeningRepository(ctx context.Context, client *ent.Client) screening.ScreeningRepository {
	return &screeningRepository{
		ctx:    ctx,
		client: client,
	}
}

func (r screeningRepository) FindByID(screeningID screening.ScreeningID) (*screening.Screening, error) {
	s, err := r.client.Screening.Get(r.ctx, string(screeningID))
	if err != nil {
		return nil, err
	}
	return reconstruct(s), nil
}

func reconstruct(s *ent.Screening) *screening.Screening {
	return &screening.Screening{
		ScreeningID:           screening.ScreeningID(s.ID),
		Status:                screening.ScreeningStatus(s.ScreeningStatus),
		ApplyDate:             s.ApplyDate,
		ApplicantEmailAddress: screening.EmailAddress(s.ApplicantEmailAddress),
	}
}

func (r screeningRepository) Insert(s *screening.Screening) error {
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

func (r screeningRepository) Update(screening *screening.Screening) error {
	// 更新処理
	return nil
}
