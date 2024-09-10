package interfaces

import (
	"context"
	"go-belajar-mock/internal/service/dto"
)

type IProductService interface {
	CreateProduct(ctx context.Context, request *dto.CreateProductRequest) error
	GetProduct(ctx context.Context, id int) (*dto.GetProductResponse, error)
}
