package screening

import (
	"context"
	"ddd-sample/ent"
)

type repository struct {
	client *ent.Client
}

type Repository interface {
	Create(ctx context.Context, screening *Screening) (*Screening, error)
}

func NewRepository(c *ent.Client) Repository {
	return &repository{client: c}
}

var _ Repository = (*MockRepository)(nil)

type MockRepository struct {
	CreateFunc func(ctx context.Context, screening *Screening) (*Screening, error)
}

func (m *MockRepository) Create(ctx context.Context, screening *Screening) (*Screening, error) {
	return m.CreateFunc(ctx, screening)
}
