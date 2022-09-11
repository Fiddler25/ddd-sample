package domain

import (
	"context"
)

type ScreeningRepository interface {
	FindByID(ctx context.Context, screeningId ScreeningID) (*Screening, error)
	Insert(ctx context.Context, screening *Screening) error
	Update(ctx context.Context, screening *Screening) error
}
