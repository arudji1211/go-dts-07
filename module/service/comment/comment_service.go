package comment

import (
	"github.com/aruji1211/myGram/module/model/comment"
	"github.com/gin-gonic/gin"
)

type CommentService interface {
	GetAll(ctx *gin.Context) ([]comment.Comment, error)
	GetById(ctx *gin.Context, idComment int) (comment.Comment, error)
	Create(ctx *gin.Context, comIn comment.Comment) (comOut comment.Comment,err error) 
	Update(ctx *gin.Context, comIn comment.Comment, idUser int) (comOut comment.Comment,err error) 
	Delete(ctx *gin.Context, idComment int, idUser int) error
}
