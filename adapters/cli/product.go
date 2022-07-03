package cli

import (
	"fmt"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/application"
)

func Run(service application.IProductService, action string, productId string, productName string, price float64) (string, error) {
	result := ""
	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product id %s was created with name %s, price %f and status %s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		response, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s was enabled", response.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		response, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s was disabled", response.GetName())
	default:
		response, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product id: %s\nName: %s\nPrice: %f\nStatus: %s", response.GetId(), response.GetName(), response.GetPrice(), response.GetStatus())
	}
	return result, nil
}
