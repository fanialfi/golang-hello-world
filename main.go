package main

import "fmt"

func tes(){
  fmt.Println("Saya fani")
}


// function yang secara otomatis akan dipanggil atau di eksekusi ketika program di jalankan
func main()  {
  fmt.Println("Hello World","selamat belajar golang")
  
  tes() // menjalankan function tes() ketika function main() di jalankan
}

// tes() // akan muncul pesan error dengan pesan syntax error : non-declaration statement outside function body
