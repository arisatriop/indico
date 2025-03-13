package repository

import (
	"fmt"
	"indico-technical-test/app/entity/request"
	"indico-technical-test/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := &LocationRepository{
		DB: &database.Postgres{GDB: gormDB},
	}

	location := &request.LocationRequest{
		Name:     "Location 1",
		Capacity: 100,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "locations"`).
		WithArgs(location.Name, location.Capacity).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create(location)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindAll(t *testing.T) {

	gotenv.Load("../../.env.test")

	err := database.NewPostgres().Set()
	if err != nil {
		t.Error(err)
	}

	db := database.Postgres{}
	pg := db.Get()

	repo := NewLocationRepository(pg)
	location, err := repo.FindAll()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(location)
}
