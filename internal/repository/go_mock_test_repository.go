package repository

import (
	"context"
	"go-belajar-mock/internal/entity"
	"go-belajar-mock/internal/repository/interfaces"
	"gorm.io/gorm"
)

type goMockTestRepository struct {
	db *gorm.DB
}

// NewGoMockTestRepository is function to create new instance goMocktestRepository
func NewGoMockTestRepository(db *gorm.DB) interfaces.IGoMockTestRepository {
	return &goMockTestRepository{db: db}
}

func (g *goMockTestRepository) Create(ctx context.Context, tx *gorm.DB, input *entity.GoMocktest) (*entity.GoMocktest, error) {
	if err := g.db.WithContext(ctx).Model(&entity.GoMocktest{}).Create(input).Error; err != nil {
		return nil, err
	}

	// success insert
	return input, nil
}

func (g *goMockTestRepository) GetByID(ctx context.Context, id int) (*entity.GoMocktest, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goMockTestRepository) GetAll(ctx context.Context) ([]entity.GoMocktest, error) {
	//TODO implement me
	panic("implement me")
}
