package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-mysql/model"
	"strconv"
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

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int) (model.Comment, error) {
	query := "SELECT id, email, comment FROM comment WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	comment := model.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next(){
		// data nya ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		// datanya tidak ada
		return comment, errors.New("<404> Id " + strconv.Itoa(id) + " Not Found.")
	}


}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]model.Comment, error) {
	query := "SELECT id, email, comment FROM comment"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		comment := model.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
	
}
