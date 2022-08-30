package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/ent/schema"
	"ddd-sample/ent/schema/property"
	"ddd-sample/src/domain/screening"
	"strconv"
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

func (r screeningRepository) FindByID(screeningId screening.ScreeningID) (*screening.Screening, error) {
	return nil, nil
}

func (r screeningRepository) Insert(s *screening.Screening) error {
	screeningID, _ := strconv.Atoi(string(s.ScreeningID))
	if _, err := r.client.Screening.
		Create().
		SetID(schema.ScreeningID(screeningID)).
		SetScreeningStatus(property.ScreeningStatus(s.Status)).
		SetNillableApplyDate(s.ApplyDate).
		SetApplicantEmailAddress(string(s.ApplicantEmailAddress)).
		Save(r.ctx); err != nil {
		return err
	}
	return nil
}

func (r screeningRepository) Update(screening *screening.Screening) error {
	return nil
}
