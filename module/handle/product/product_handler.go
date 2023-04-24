package product

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
