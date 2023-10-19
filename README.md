1. package database
    - sudah dibuat setandarnya dalam package database.

2. database driver
    - wajib menambahkan database driver

3. membuat koneksi
    - username:password@tcp(host:port)/db_name

4. Database pooling
    - secara default ketika kita membuat koneksi di golang, golang akan memasukkan koneksi tersebut ke dalam pooling. kegunaanya adalah ketika ada sesuatu yang membutuhkan koneksi tidak harus membuat koneksi secara manual, tetapi langsung kirimkan saja perintah sql nya.

    - jadi kita bisa membatasi koneksi ke database. batas minimal dan juga maksimal. jadi nantinya jika aplikasi pertama kali di jalankan golang akan membuatkan minimal koneksi yang sudah tersetting, dan ketika trafic nya naik maka golang akan membuat koneksi lagi hingga batas maksimum.

    - setting database pooling ini wajib dilakukan ketika membuat database digolang.

5. Eksekusi perintah SQL
    - gunakan ExecContext() jika kita tidak membutuhkan data kembalian dari databasenya seperti (select), select membutuhkan data jadi jangan gunakan ExecContext(). 

6. Eksekusi perintah SELECT SQL
    - gunakan QueryContext() untuk query yang mengembalikan hasil seperti Select.

7. Tipe data column
    - mengambil data dari database lalu memasukkan ke tipe data yang sudah dijabarkan. 

    - jika kita membutuhkan tipe data date time sebenarnya kita bisa saja mengkonversinya ke string lalu ke date, tetapi itu akan capek. tambahkan saja parseTime=true pada koneksinya untuk mengkonversi secara otomatis untuk tipe datetime. 

8. Auto Increment
    - gunakan result pada query insert atau query yang tidak mengembalikan hasil, ini kan mengembalikan id yang ada pada tabel yang bersifat auto increment.

9. Prepare Statement
    - digunakan ketika kita membutuhkan esekusi query berkali kali dengan database koneksi yang sama dan dengan query yang sama, hanya berbeda parameternya.

10. Transaction

