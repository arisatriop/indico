package repository

import (
	"fmt"
	"indico-technical-test/database"
	"testing"

	"github.com/subosito/gotenv"
)

func TestUserRepository(t *testing.T) {
	t.Run("FindAll", func(t *testing.T) {
		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewUserRepository(pg)
		user, err := repo.FindAll()
		if err != nil {
			t.Error(err)
		}

		fmt.Println(user)
	})

	t.Run("FindMe", func(t *testing.T) {
		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewUserRepository(pg)
		users, err := repo.FindAll()
		if err != nil {
			t.Error(err)
		}

		user, err := repo.FindMe(users[0].Token)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(user)
	})
}
