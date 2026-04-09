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

