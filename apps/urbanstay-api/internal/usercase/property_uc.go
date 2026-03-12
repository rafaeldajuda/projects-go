package usercase

import (
	"context"
	"errors"
	"urbanstay-api/internal/domain"
	"urbanstay-api/internal/domain/entity"
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
	if p.Name == "" {
		return errors.New("o nome do imóvel é obrigatório")
	}
	if p.PricePerNight <= 0 {
		return errors.New("o preço por noite deve ser maior que zero")
	}

	// Persistencia
	// O Usecase não sabe se o 'repo' salva no MySQL ou na Memória
	err := uc.repo.AddProperty(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PropertyUseCase) ExecuteList(ctx context.Context) []*entity.Property {
	return uc.repo.ListProperties(ctx)
}
