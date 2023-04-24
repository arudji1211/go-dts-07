package product

import (
	"github.com/arudji1211/go-dts-07/module/model/product"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAll(ctx *gin.Context) ([]product.Product, error)
	GetById(ctx *gin.Context, idPhoto int) (product.Product, error)
	Create(ctx *gin.Context, productIn product.Product) (productOut product.Product, err error)
	Update(ctx *gin.Context, productIn product.Product, idUser int) (productOut product.Product, err error)
	Delete(ctx *gin.Context, idPhoto int, idUser int) error
}
