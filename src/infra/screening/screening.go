package screening

import (
	"context"
	"ddd-sample/ent"
	"ddd-sample/ent/schema"
	"ddd-sample/ent/schema/property"
	"ddd-sample/sdk/convert"
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

func (r screeningRepository) FindByID(id screening.ScreeningID) (*screening.Screening, error) {
	screeningID := convert.StrToInt[schema.ScreeningID](string(id))
	s, err := r.client.Screening.Get(r.ctx, screeningID)
	if err != nil {
		return nil, err
	}
	return reconstruct(s), nil
}

func reconstruct(s *ent.Screening) *screening.Screening {
	return &screening.Screening{
		ScreeningID:           screening.ScreeningID(strconv.Itoa(int(s.ID))),
		Status:                screening.ScreeningStatus(s.ScreeningStatus),
		ApplyDate:             s.ApplyDate,
		ApplicantEmailAddress: screening.EmailAddress(s.ApplicantEmailAddress),
	}
}

func (r screeningRepository) Insert(s *screening.Screening) error {
	screeningID := convert.StrToInt[schema.ScreeningID](string(s.ScreeningID))
	if _, err := r.client.Screening.
		Create().
		SetID(screeningID).
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
