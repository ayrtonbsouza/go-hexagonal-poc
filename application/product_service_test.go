package application_test

import (
	"testing"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
	mock_application "github.com/ayrtonbsouza/hexagonal-architecture-poc/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, _ := service.Get("abc")

	require.Equal(t, product, result)
}
