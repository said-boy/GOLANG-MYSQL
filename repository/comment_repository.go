package repository

import (
	"context"
	"golang-mysql/model"
)

type ICommentRepository interface{
	Insert(ctx context.Context, comment model.Comment) (model.Comment, error)
}