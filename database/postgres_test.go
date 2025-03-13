package database

import (
	"testing"

	"github.com/subosito/gotenv"
)

func TestSetConnection(t *testing.T) {

	gotenv.Load("../.env.test")
	t.Run("Test Set Connection", func(t *testing.T) {
		err := NewPostgres().Set()
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}
	})
}

func TestGetConnection(t *testing.T) {

	gotenv.Load("../.env.test")

	t.Run("Test Get Connection", func(t *testing.T) {
		_ = NewPostgres().Set()

		pg := Postgres{}
		db := pg.Get()
		if db.DB == nil {
			t.Error("Test failed: db is nil")
		}
		if db.GDB == nil {
			t.Error("Test failed: gdb is nil")
		}
	})
}
