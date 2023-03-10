package screening

import (
	"context"
	"ddd-sample/ent/schema/property"
)

func (r *repository) Insert(ctx context.Context, s *Screening) error {
	if _, err := r.client.Screening.
		Create().
		SetID(string(s.ID)).
		SetScreeningStatus(property.ScreeningStatus(s.Status)).
		SetNillableApplyDate(s.ApplyDate).
		SetApplicantEmailAddress(s.ApplicantEmailAddress).
		Save(ctx); err != nil {
		return err
	}
	return nil
}
