package golang_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
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

// 7. Tipe data column
func TestSelectTipeData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// letak column pada query harus sama dengan letak pada rows.Scan()
	query := "SELECT id, name, tgl_lahir, menikah, hobi FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next(){
		var id int
		var name string
		
		// ketika kita membutuhkan data dengan tipe datetime 
		// kita harus menambahkan parseTime=true pada koneksinya.
		// maka tipe datetime akan secara otomatis dikonversi. 
		var tgl_lahir time.Time
		var menikah bool
		var hobi sql.NullString // untuk column yang bisa null
		/* sqlNull, dll akan mengembalikan 2 yaitu : String -> value dan Valid -> boolean */

		err := rows.Scan(&id, &name, &tgl_lahir, &menikah, &hobi)
		if err != nil {
			panic(err)
		}

		fmt.Println("===========================")
		fmt.Println("id : ", id)
		fmt.Println("name : ", name)
		fmt.Println("tgl_lahir : ", tgl_lahir)
		fmt.Println("menikah : ", menikah)
		if hobi.Valid {
			fmt.Println("hobi : ", hobi)
		}
		fmt.Println("===========================")

	}

	defer rows.Close()

}

// 8. Auto Increment
func TestGetAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO customer(name, tgl_lahir, menikah, hobi) VALUES (?, ?, ?, ?) "
	result, err := db.ExecContext(ctx, query, "budi", time.Now(), false, "membaca")
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Id data yang ditambahkan adalah ", insertId)

}

// 10. Prepare Statement
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i++ {
		email := "said" + strconv.Itoa(i) + "@gmail.com"
		comment := "Halo, Ini comment ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		fmt.Println("Succes insert data id ", insertId)
	}

	defer stmt.Close()

}


