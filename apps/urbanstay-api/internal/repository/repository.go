package repository

import (
	"context"
	"urbanstay-api/internal/domain/entity"
)

type PropertyRepository interface {
	AddProperty(ctx context.Context, p *entity.Property) error
	ListProperties(ctx context.Context) ([]*entity.Property, error)
}
