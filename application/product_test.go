package application_test

import (
	"testing"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
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
