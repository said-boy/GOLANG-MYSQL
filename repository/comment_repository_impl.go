package repository

import (
	"context"
	"database/sql"
	"golang-mysql/model"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) ICommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment model.Comment) (model.Comment, error) {
	query := "INSERT INTO comment(email, comment) VALUES(?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(insertId)

	return comment, nil

}