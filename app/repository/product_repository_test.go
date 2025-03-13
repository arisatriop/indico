package repository

import (
	"fmt"
	"indico-technical-test/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestProductRepository(t *testing.T) {

	t.Run("Delete", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		gormDB, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})

		repo := &ProductRepository{
			DB: &database.Postgres{GDB: gormDB},
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "products"`).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.Delete(1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("FindAll", func(t *testing.T) {
		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewProductRepository(pg)
		product, err := repo.FindAll()
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("len product: %d\n", len(product))
	})

	t.Run("FindByID", func(t *testing.T) {
		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewProductRepository(pg)
		product, err := repo.FindByID(1)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("product: %v\n", product)
	})
}
