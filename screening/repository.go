package screening

import (
	"context"
	"ddd-sample/ent"
)

type repository struct {
	client *ent.Client
}

type Repository interface {
	Insert(ctx context.Context, screening *Screening) error
}

func NewRepository(c *ent.Client) Repository {
	return &repository{client: c}
}
