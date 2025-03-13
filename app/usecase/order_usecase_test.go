package usecase

import (
	"fmt"
	"testing"
)

func TestGenerateOrderNumber(t *testing.T) {

	orderUsecase := OrderUsecase{}

	t.Run("1", func(t *testing.T) {
		orderNumber := orderUsecase.GenerateOrderNumber()
		if orderNumber == "" {
			t.Errorf("Expected not empty, got empty")
		}
	})

	t.Run("2", func(t *testing.T) {
		orderNumber := orderUsecase.GenerateOrderNumber()
		fmt.Println(orderNumber)
	})

	t.Run("3", func(t *testing.T) {
		orderNumber := orderUsecase.GenerateOrderNumber()
		if len(orderNumber) != 22 {
			t.Errorf("Expected length 22, got %d", len(orderNumber))
		}
	})
}
