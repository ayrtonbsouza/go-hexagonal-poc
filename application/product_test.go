package application_test

import (
	"testing"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "test"
	product.Status = application.DISABLED
	product.Price = 100

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price should be greater than 0", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 100

	err = product.Disable()
	require.Equal(t, "price should be 0", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "test"
	product.Status = application.DISABLED
	product.Price = 100

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "status should be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "price should be greater or equal than 0", err.Error())

	product.Price = 0
	_, err = product.IsValid()
	require.Nil(t, err)
}
