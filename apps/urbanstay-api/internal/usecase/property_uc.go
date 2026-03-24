package usecase

import (
	"context"
	"time"
	"urbanstay-api/internal/domain"
	"urbanstay-api/internal/domain/entity"

	"github.com/google/uuid"
)

type PropertyUseCase struct {
	repo domain.PropertyRepository
}

func NewPropertyUseCase(repo domain.PropertyRepository) *PropertyUseCase {
	return &PropertyUseCase{
		repo: repo,
	}
}

func (uc *PropertyUseCase) ExecuteCreate(ctx context.Context, p *entity.Property) error {
	// Regra de negócio
	if err := p.Validate(); err != nil {
		return err
	}

	// Gen ID
	p.ID = uuid.NewString()

	// Add created_at
	p.CreatedAt = time.Now()

	// Persistencia
	// O Usecase não sabe se o 'repo' salva no MySQL ou na Memória
	err := uc.repo.AddProperty(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PropertyUseCase) ExecuteList(ctx context.Context) ([]*entity.Property, error) {
	return uc.repo.ListProperties(ctx)
}
