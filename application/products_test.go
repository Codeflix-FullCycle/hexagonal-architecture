package application_test

import (
	"testing"

	application "github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProducts_Enable(t *testing.T) {
	product := application.Products{}

	product.Name = "Coffee"
	product.Price = 0
	product.Status = application.DISABLED

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Enable()
	require.Equal(t, "the price most be greater than zero to enable the product", err.Error())

}

func TestProducts_Disable(t *testing.T) {
	product := application.Products{}

	product.Name = "Coffee"
	product.Price = 0
	product.Status = application.ENABLED

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.Error(t, err)

}

func TestProducts_IsValid(t *testing.T) {
	product := application.Products{}
	product.ID = uuid.NewV4().String()
	product.Name = "Coffee"
	product.Price = 10
	product.Status = application.DISABLED

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "status_test"

	_, err = product.IsValid()

	require.Equal(t, err.Error(), "the status must be enable or disable")

	product.Status = application.DISABLED
	product.Price = -10

	_, err = product.IsValid()

	require.Equal(t, err.Error(), "the price must greater or equal zero")

	product.ID = "123124"

	_, err = product.IsValid()

	require.Error(t, err)
}
