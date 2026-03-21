package memory

import (
	"context"
	"urbanstay-api/internal/domain/entity"
)

type MemoryRepository struct {
	Properties []*entity.Property
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (m *MemoryRepository) AddProperty(ctx context.Context, p *entity.Property) error {
	m.Properties = append(m.Properties, p)
	return nil
}

func (m *MemoryRepository) ListProperties(ctx context.Context) []*entity.Property {
	list := make([]*entity.Property, len(m.Properties))
	for k, v := range m.Properties {
		if v != nil {
			cp := *v
			list[k] = &cp
		}
	}
	return list
}
