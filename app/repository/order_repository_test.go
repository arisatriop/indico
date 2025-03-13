package repository

import (
	"fmt"
	"indico-technical-test/database"
	"testing"

	"github.com/subosito/gotenv"
)

func TestOrderRepository(t *testing.T) {
	t.Run("FindAll", func(t *testing.T) {
		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewOrderRepository(pg)
		order, err := repo.FindAll()
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("Order: %+v\n", len(order))
	})

	t.Run("FindByID", func(t *testing.T) {
		gotenv.Load("../../.env.test")

		err := database.NewPostgres().Set()
		if err != nil {
			t.Error(err)
		}

		db := database.Postgres{}
		pg := db.Get()

		repo := NewOrderRepository(pg)
		order, err := repo.FindByID(33)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("Order: %v\n", order)
	})
}
