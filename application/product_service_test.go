package application_test

import (
	"testing"

	"github.com/Codeflix-FullCycle/hexagonal-architecture/application"
	mock_application "github.com/Codeflix-FullCycle/hexagonal-architecture/application/mocks"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductsInterface(ctrl)
	fakePersistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	// mock get method
	fakePersistence.EXPECT().Get(gomock.Any()).Return(fakeProduct, nil).AnyTimes()

	productService := application.ProductService{
		ProductPersistence: fakePersistence,
	}

	result, err := productService.Get("123")

	require.Nil(t, err)
	require.Equal(t, result, fakeProduct)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductsInterface(ctrl)
	fakePersistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	fakePersistence.EXPECT().Save(gomock.Any()).Return(fakeProduct, nil).AnyTimes()

	service := application.ProductService{
		ProductPersistence: fakePersistence,
	}

	result, err := service.Create("first product", 1)

	require.Nil(t, err)
	require.Equal(t, result, fakeProduct)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductsInterface(ctrl)
	fakePersistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	fakeProduct.EXPECT().Enable().Return(nil)
	// mock get method
	fakePersistence.EXPECT().Save(gomock.Any()).Return(fakeProduct, nil).AnyTimes()

	productService := application.ProductService{
		ProductPersistence: fakePersistence,
	}

	result, err := productService.Enable(fakeProduct)

	require.Nil(t, err)
	require.Equal(t, result, fakeProduct)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductsInterface(ctrl)
	fakePersistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	fakeProduct.EXPECT().Disable().Return(nil)

	// mock get method
	fakePersistence.EXPECT().Save(gomock.Any()).Return(fakeProduct, nil).AnyTimes()

	productService := application.ProductService{
		ProductPersistence: fakePersistence,
	}

	result, err := productService.Disable(fakeProduct)

	require.Nil(t, err)
	require.Equal(t, result, fakeProduct)
}
