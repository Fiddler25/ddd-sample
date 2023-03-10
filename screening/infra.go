package screening

import (
	"context"
	"ddd-sample/ent/schema/property"
)

func (r *repository) Create(ctx context.Context, s *Screening) (*Screening, error) {
	e, err := r.client.Screening.
		Create().
		SetID(string(s.ID)).
		SetScreeningStatus(property.ScreeningStatus(s.Status)).
		SetNillableApplyDate(s.ApplyDate).
		SetApplicantEmailAddress(s.ApplicantEmailAddress).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &Screening{
		ID:                    ID(e.ID),
		Status:                Status(e.ScreeningStatus),
		ApplyDate:             e.ApplyDate,
		ApplicantEmailAddress: e.ApplicantEmailAddress,
	}, nil
}
