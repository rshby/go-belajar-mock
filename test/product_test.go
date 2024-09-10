package test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-belajar-mock/internal/service"
	"go-belajar-mock/internal/service/dto"
	mock_interfaces "go-belajar-mock/test/mock/repository"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	dbMock.ExpectQuery("SELECT VERSION()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.23"))

	dbMysql, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), nil)
	if err != nil {
		t.Fatal(err)
	}

	repo := mock_interfaces.NewMockIGoMockTestRepository(mockCtrl)
	srv := service.NewProductService(dbMysql, repo)

	//t.Run("test success create", func(t *testing.T) {
	dbMock.ExpectBegin()
	dbMock.ExpectQuery("INSERT INTO go_mock_test (identity_number, full_name) VALUES (?, ?) RETURNING id").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	//dbMock.ExpectCommit()

	repo.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
	err = srv.CreateProduct(context.Background(), &dto.CreateProductRequest{
		IdentityNumber: "123",
		FullName:       "Reo Sahobby",
	})
	assert.Nil(t, err)
	//})

	/*
		t.Run("test failed to create", func(t *testing.T) {
			errorMessage := "error failed to insert new data"
			repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New(errorMessage))
			err := srv.CreateProduct(context.Background(), &dto.CreateProductRequest{
				IdentityNumber: "123",
				FullName:       "Reo Sahobby",
			})

			assert.NotNil(t, err)
			assert.EqualError(t, err, errorMessage)
		})
	*/
}
