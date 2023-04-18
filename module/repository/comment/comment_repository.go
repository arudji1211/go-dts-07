package comment

import (
	"context"

	"github.com/aruji1211/myGram/module/model/comment"
)

type CommentRepository interface {
	GetAll(ctx context.Context) ([]comment.Comment, error)
	GetById(ctx context.Context, idComment int) (comment.Comment, error)
	Create(ctx context.Context, comIn comment.Comment) (comOut comment.Comment,err error)
	Update(ctx context.Context, comIn comment.Comment) (comOut comment.Comment,err error)
	Delete(ctx context.Context, idComment int) error
}
