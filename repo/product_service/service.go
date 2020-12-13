package product_service

import (
	"go_id/services/config"
	"go_id/services/repo/product_service/schema"
)

type IProductService interface {
	Find(productId int64) (product *schema.Product, err error)
	Add(product schema.Product) bool
}

type productService struct {
	url string
}

func NewProduct() IProductService {
	var ps productService
	ps.url = config.ServiceConfig().ProductSrvAddr
	return &ps
}

func (ps *productService) Find(productId int64) (product *schema.Product, err error) {
	return product, err
}

func (ps *productService) Add(product schema.Product) bool {
	return true
}
