package golang_mysql

import (
	"database/sql"
	"time"
)

// 4. Database pooling
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/golang_database")
	if err != nil {
		panic(err)
	}
	// golang akan membuat 10 koneksi database ketika pertama kali dijalankan (batas minimal koneksi)
	db.SetMaxIdleConns(10)

	// maksimal koneksi yang diizinkan dibuat oleh golang
	// jika ada membutuhkan koneksi ketika sudah batas maksimal
	// maka akan disuruh mengantri
	db.SetMaxOpenConns(100)

	// waktu jeda sebelum koneksi ditutup jika koneksi sudah tidak lagi digunakan
	db.SetConnMaxIdleTime(5 * time.Minute)

	// waktu batas maksimal penggunaan koneksi
	// jika ada koneksi yang di gunakan lebih dari 1 jam
	// maka koneksi tersebut akan dihapus 
	// dan diganti dengan koneksi yang baru
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
