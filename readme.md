# ?

- keyword `package`

setiap file program harus memiliki `package` dan setiap folder project harus memiliki minimal 1 package main. File yang berpackage main, akan di eksekusi pertama kali ketika program di jalankan.

```go
package <nama-package>
package main
```

- keyword `import`

digunakan untuk meng-`import` atau memasukkan package lain ke dalam file program. Package fmt merupakan salah satu package bawaan yang di sediakan oleh GO, isinya banyak fungsi untuk keperluan dengan I/O yang berhubungan dengan text.

```go
import <nama-package>
import "fmt"
```

- `main` function

dalam sebuah project harus ada function `main` yang tersimpan didalam file yang package-nya `main`. function `main()` adalah function yang pertama kali di panggil ketika program di eksekusi.

```go
func main(){
    // do something here
}
```

- function `fmt.Println()`

digunakan untuk memunculkan teks ke stdout, skema penulisannya bisa dilihat sebagai berikut

```go
fmt.Println("<isi-pesan>")
fmt.Println("Hello World")
```

function tersebut berada di package `fmt` jadi untuk penggunaannya perlu melakukan import package `fmt` terlebih dahulu, function `fmt.Println()` dapat menampung parameter yang jumplahnya tidak terbatas, semua parameter akan ditampilkan dengan menggunakan spasi sebagai pemisahnya.

```go
fmt.Println("Hello","World","How","Are","You")
```
