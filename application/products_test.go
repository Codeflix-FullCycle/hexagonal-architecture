package application_test

import (
	"testing"

	application "github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
)

func TestProducts_Enable(t *testing.T) {
	product := application.Products{}

	product.Name = "Coffee"
	product.Price = 10
	product.Status = application.DISABLED

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price most be greater than zero to enable the product", err.Error())

}
