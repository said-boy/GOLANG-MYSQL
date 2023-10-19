package repository

import (
	"context"
	"golang-mysql/model"
)

type ICommentRepository interface{
	Insert(ctx context.Context, comment model.Comment) (model.Comment, error)
	FindById(ctx context.Context, id int) (model.Comment, error)
	FindAll(ctx context.Context) ([]model.Comment, error)
}