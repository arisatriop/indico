package repository

import (
	"indico-technical-test/app/entity"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAuthRepository(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := &AuthRepository{
		DB: &database.Postgres{GDB: gormDB},
	}

	user := &entity.User{
		Name:     "Staff Updated",
		Username: "staff",
		Token:    "test",
		Password: "test",
		Roles:    "staff",
	}

	t.Run("Register", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).
			WithArgs(user.Name, user.Username, user.Password, user.Token, user.Roles).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err = repo.Register(user)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "users"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.Update(user)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

	t.Run("Login", func(t *testing.T) {

		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewAuthRepository(pg)
		user, err := repo.Login(&request.LoginRequest{
			Username: "staff",
		})
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Username, "staff")
	})

}
