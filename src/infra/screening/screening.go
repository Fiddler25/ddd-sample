package screening

import (
	"context"
	"ddd-sample/ent"
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

func (r screeningRepository) FindByID(screeningId screening.ScreeningID) (*screening.Screening, error) {
	return nil, nil
}

func (r screeningRepository) Insert(screening *screening.Screening) error {
	return nil
}

func (r screeningRepository) Update(screening *screening.Screening) error {
	return nil
}
