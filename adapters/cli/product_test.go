package cli_test

import (
	"fmt"
	"testing"

	"github.com/Codeflix-FullCycle/hexagonal-architecture/adapters/cli"
	mock_application "github.com/Codeflix-FullCycle/hexagonal-architecture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Test"
	productId := "123"
	productStatus := "disable"
	productPrice := 10.0

	product_mock := mock_application.NewMockProductsInterface(ctrl)

	product_mock.EXPECT().GetID().Return(productId).AnyTimes()
	product_mock.EXPECT().GetName().Return(productName).AnyTimes()
	product_mock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	product_mock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	service_mock := mock_application.NewMockProductServiceInterface(ctrl)

	service_mock.EXPECT().Create(productName, productPrice).Return(product_mock, nil).AnyTimes()
	service_mock.EXPECT().Get(productId).Return(product_mock, nil).AnyTimes()
	service_mock.EXPECT().Enable(gomock.Any()).Return(product_mock, nil).AnyTimes()
	service_mock.EXPECT().Disable(gomock.Any()).Return(product_mock, nil).AnyTimes()

	expected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service_mock, "create", productId, productName, productPrice)
	require.Empty(t, err)
	require.Equal(t, result, expected)

	expected = fmt.Sprintf("Product %s has been enabled.", productName)

	result, err = cli.Run(service_mock, "enable", productId, productName, productPrice)
	require.Empty(t, err)
	require.Equal(t, result, expected)

	expected = fmt.Sprintf("Product %s has been disable.", productName)

	result, err = cli.Run(service_mock, "disable", productId, productName, productPrice)
	require.Empty(t, err)
	require.Equal(t, result, expected)

	expected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)

	result, err = cli.Run(service_mock, "get", productId, productName, productPrice)
	require.Empty(t, err)
	require.Equal(t, result, expected)

}
