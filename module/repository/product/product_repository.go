package product

import (
	"context"

	"github.com/arudji1211/go-dts-07/module/model/product"
)

type ProductRepository interface {
	GetAll(ctx context.Context) (productOut []product.Product, err error)
	GetAllByUserId(ctx context.Context, userId int) (productOut []product.Product, err error)
	GetById(ctx context.Context, idProduct int) (productOut product.Product, err error)
	Create(ctx context.Context, productIn product.Product) (productOut product.Product, err error)
	Update(ctx context.Context, productIn product.Product) (productOut product.Product, err error)
	Delete(ctx context.Context, idProduct int) (err error)
}
