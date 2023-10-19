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