package main

import(
	"fmt"
)

func add(a int,  b int){
	fmt.Printf("Hasil penjumlahan %d dan %d adalah : %d", a, b,(a+b))
}

func main() {
	add(1,2)
}