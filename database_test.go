package golang_mysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // menggunakan _ karena hanya butuh function init() saja.
)

func TestEmpty(t *testing.T) {

}

// 3. Membuat koneksi
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()	
}




