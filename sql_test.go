package golang_mysql

import (
	"context"
	"fmt"
	"testing"
)

// 5. Eksekusi perintah SQL
func TestInsertSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(name) VALUES (?)" // untuk menghindari sql-injection

	// ExecContext -> digunakan jika kita mengirim 
	// query yang tidak return data seperti select.
	_ , err := db.ExecContext(ctx, query,"said")
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert data to database")
}

// 6. Esekusi Perintah SELECT SQL 
func TestSelectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id : ", id, " name : ", name)

	}

	defer rows.Close()

}


