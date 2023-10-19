package repository

import (
	"context"
	"fmt"
	golang_mysql "golang-mysql"
	"golang-mysql/model"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestInsertComment(t *testing.T) {
	db := golang_mysql.GetConnection()
	comment := NewCommentRepository(db)

	ctx := context.Background()
	modelComment := model.Comment{
		Email: "Said@repository.com",
		Comment: "Halo, ini repository",
	}

	result, err := comment.Insert(ctx, modelComment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	comment := NewCommentRepository(golang_mysql.GetConnection())
	result, err := comment.FindById(context.Background(), 20)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	comment := NewCommentRepository(golang_mysql.GetConnection())
	comments, err := comment.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	} 
}