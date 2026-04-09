# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[MUHAMMADTETUKOKEMALPASHA] - [109082500181]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go

package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	fmt.Print("n  ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d  ", i)
	}
	fmt.Println() 
                                     
	fmt.Print("Sn ")
	for i := 0; i <= n; i++ {
		
		fmt.Printf("%d  ", fibonacci(i))
	}
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_MODUL5/blob/main/MODUL5/109082500124_MUHAMMADFAIZMAULANA-main/modul5/output/out-put%20soal1.png)
[penjelasan] :Program ini menghitung dan menampilkan deret Fibonacci dari indeks 0 sampai nilai n yang dimasukkan pengguna. Fungsi fibonacci(n) dibuat secara rekursif, dengan kondisi dasar n=0 mengembalikan 0 dan n=1 mengembalikan 1, sedangkan untuk nilai lainnya dihitung dari penjumlahan dua suku sebelumnya, yaitu fibonacci(n-1) + fibonacci(n-2). Di fungsi main, program meminta input n, lalu mencetak dua baris output: baris pertama berisi indeks n (0 hingga n), dan baris kedua berisi nilai setiap suku Fibonacci (Sn) yang sesuai dengan masing-masing indeks tersebut.


## Unguided 

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 0)
}

func bintang(n, i int) {
	if i != n {
		baris(i,0)
		fmt.Println("")
		bintang(n, i+1)
	}else{
		baris(i,0)
	}
}

func baris(i, a int) {
	if a != i {
		fmt.Print("*")
		baris(i, a+1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_MODUL5/blob/main/MODUL5/109082500124_MUHAMMADFAIZMAULANA-main/modul5/output/out-put%20soal2.png)
[penjelasan] : Program ini membaca sebuah bilangan n lalu mencetak pola segitiga bintang secara rekursif, dimulai dari baris kosong hingga baris dengan jumlah bintang terbanyak. Fungsi bintang(n, i) mengatur perpindahan antarbaris dengan menaikkan nilai i dari 0 sampai n, sedangkan fungsi baris(i, a) mencetak karakter * sebanyak i kali pada tiap baris juga dengan rekursi. Saat i belum sama dengan n, program mencetak baris lalu lanjut ke baris berikutnya; ketika mencapai kondisi akhir, program mencetak baris terakhir dan berhenti, sehingga terbentuk pola bintang bertingkat tanpa menggunakan perulangan biasa untuk proses utamanya.
## Unguided 

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func cariFaktor(n, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
		cariFaktor(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	fmt.Print("Faktor: ")
	cariFaktor(n, 1)
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_MODUL5/blob/main/MODUL5/109082500124_MUHAMMADFAIZMAULANA-main/modul5/output/out-put%20soal3.png)
[penjelasan] :berfungsi untuk menampilkan semua faktor dari sebuah bilangan bulat n yang dimasukkan pengguna. Setelah input dibaca di fungsi main, program memanggil fungsi rekursif cariFaktor(n, 1) yang memeriksa setiap angka i mulai dari 1 hingga n; jika n habis dibagi i (n%i == 0), maka i dicetak sebagai faktor. Proses ini berulang secara rekursif dengan menaikkan i satu per satu sampai melewati n, sehingga seluruh faktor bilangan dapat ditampilkan tanpa menggunakan perulangan biasa.

## Unguided 

### 4. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var n,i int = 0,1
	fmt.Scan(&n)
	turun(n,i)
	naik(n,i)
}

func turun(n, i int) {
	if n != i{
	fmt.Print(n)
	fmt.Print(" ")
	turun(n-1,i)
	}
}

func naik(n,i int){
	if i <= n{
		fmt.Print(i)
		fmt.Print(" ")
		naik(n,i+1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_MODUL5/blob/main/MODUL5/109082500124_MUHAMMADFAIZMAULANA-main/modul5/output/out-put%20soal4.png)
[penjelasan] :membaca sebuah bilangan n dari input, lalu menampilkan dua pola angka menggunakan rekursi: pertama urutan menurun dari n sampai 2 melalui fungsi turun, kemudian urutan menaik dari 1 sampai n melalui fungsi naik. Fungsi turun(n, i) terus mencetak nilai n dan memanggil dirinya sendiri dengan n-1 selama n tidak sama dengan i (dengan i bernilai awal 1), sedangkan fungsi naik(n, i) mencetak i lalu menaikkan nilainya satu per satu hingga i lebih besar dari n. Hasil akhirnya adalah deretan angka menurun lalu menaik yang dicetak berurutan dalam satu baris.

## Unguided 

### 5. [Soal]
#### soal1.go

```go
package main

import "fmt"

func cetakGanjil(n, i int) {
	if i <= n {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
		cetakGanjil(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	cetakGanjil(n, 1)
	fmt.Println()
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_MODUL5/blob/main/MODUL5/109082500124_MUHAMMADFAIZMAULANA-main/modul5/output/out-put%20soal5.png)
[penjelasan] :digunakan untuk menampilkan semua bilangan ganjil dari 1 sampai n, dengan n dimasukkan oleh pengguna. Setelah input dibaca di fungsi main, program memanggil fungsi rekursif cetakGanjil(n, 1) yang memeriksa nilai i mulai dari 1 hingga n; jika i bernilai ganjil (i%2 != 0), angka tersebut dicetak ke layar. Kemudian fungsi memanggil dirinya sendiri dengan i+1 sampai batas i > n tercapai, sehingga seluruh bilangan ganjil dalam rentang tersebut ditampilkan tanpa menggunakan perulangan biasa.

## Unguided 

### 6. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var x, y int
	var hasil int = 1
	fmt.Scan(&x,&y)
	pangkat(x,y,hasil)
}

func pangkat(x, y, hasil int){
	if y > 0{
		hasil = x * hasil
		pangkat(x,y-1,hasil)
	}else{
		fmt.Print(hasil)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_MODUL5/blob/main/MODUL5/109082500124_MUHAMMADFAIZMAULANA-main/modul5/output/out-put%20soal6.png)
[penjelasan] :Program ini menghitung hasil perpangkatan secara rekursif, yaitu nilai x pangkat y dari dua input bilangan bulat yang dimasukkan pengguna. Di fungsi main, program membaca x dan y lalu memanggil fungsi pangkat dengan nilai awal hasil = 1 sebagai penampung akumulasi perkalian. Pada fungsi pangkat, selama y masih lebih dari 0, nilai hasil terus dikalikan dengan x dan fungsi memanggil dirinya sendiri dengan y-1; ketika y sudah 0, proses berhenti dan nilai hasil dicetak sebagai output akhir. Dengan alur ini, operasi perpangkatan dilakukan melalui perkalian berulang menggunakan rekursi.