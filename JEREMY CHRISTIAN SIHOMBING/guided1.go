package main
import "fmt"

func main() {

 var detik, jam, menit int
 fmt.Print("Masukkan detik: ")
 fmt.Scan(&detik)
 jam = detik / 3600
 menit = (detik % 3600) / 60
 detik = detik % 60
 fmt.Printf("Hasil konversi: %d jam %d menit dan %d detik\n", jam, menit, detik)
}
