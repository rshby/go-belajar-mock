package interfaces

import (
	"context"
	"go-belajar-mock/internal/entity"
	"gorm.io/gorm"
)

type IGoMockTestRepository interface {
	Create(ctx context.Context, tx *gorm.DB, input *entity.GoMocktest) (*entity.GoMocktest, error)
	GetByID(ctx context.Context, id int) (*entity.GoMocktest, error)
	GetAll(ctx context.Context) ([]entity.GoMocktest, error)
}
