package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"go-belajar-mock/internal/entity"
	"go-belajar-mock/internal/repository/interfaces"
	"go-belajar-mock/internal/service/dto"
	interfaces2 "go-belajar-mock/internal/service/interfaces"
	"gorm.io/gorm"
)

type productService struct {
	db            *gorm.DB
	gmtRepository interfaces.IGoMockTestRepository
}

func NewProductService(db *gorm.DB, gmtRepository interfaces.IGoMockTestRepository) interfaces2.IProductService {
	return &productService{db, gmtRepository}
}

func (srv *productService) CreateProduct(ctx context.Context, request *dto.CreateProductRequest) error {
	// create input
	input := entity.GoMocktest{
		IdentityNumber: request.IdentityNumber,
		FullName:       request.FullName,
	}

	// create transaction
	tx := srv.db.Begin()
	defer tx.Rollback()

	// insert
	_, err := srv.gmtRepository.Create(ctx, tx, &input)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// commit transaction
	if err = tx.Commit().Error; err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (srv *productService) GetProduct(ctx context.Context, id int) (*dto.GetProductResponse, error) {
	//TODO implement me
	panic("implement me")
}
