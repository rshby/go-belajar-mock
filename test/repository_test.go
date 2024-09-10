package test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-belajar-mock/internal/entity"
	"go-belajar-mock/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestCreateData(t *testing.T) {
	dbMysql, s, _ := sqlmock.New()
	defer dbMysql.Close()

	mysqlDialector := mysql.New(mysql.Config{
		Conn: dbMysql,
	})

	s.ExpectQuery("SELECT VERSION()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.23"))

	db, err := gorm.Open(mysqlDialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// create instance repository
	testRepository := repository.NewGoMockTestRepository(db)

	t.Run("test success insert", func(t *testing.T) {
		s.ExpectBegin()

		query := regexp.QuoteMeta("INSERT INTO  `go_mock_test` (`created_at`,`updated_at`,`identity_number`,`full_name`)")
		s.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "123", "Reo Sahobby").WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)
		s.ExpectCommit()

		create, errInsertt := testRepository.Create(context.Background(), db, &entity.GoMocktest{
			IdentityNumber: "123",
			FullName:       "Reo Sahobby",
		})

		assert.Nil(t, errInsertt)
		assert.NotNil(t, create)
		assert.Equal(t, 1, create.ID)
	})

	t.Run("test failed insert", func(t *testing.T) {
		s.ExpectBegin()
		s.ExpectExec(regexp.QuoteMeta("INSERT INTO  `go_mock_test` (`created_at`,`updated_at`,`identity_number`,`full_name`)")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error insert"))
		s.ExpectRollback()

		result, errInsert := testRepository.Create(context.Background(), db, &entity.GoMocktest{
			IdentityNumber: "123",
			FullName:       "Reo",
		})

		assert.Error(t, errInsert)
		assert.NotNil(t, errInsert)
		assert.EqualError(t, errInsert, "error insert")
		assert.Nil(t, result)
	})
}
