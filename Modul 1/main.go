package main 

import (
	"fmt"
	"strings"
)

func main (){
	n1 := 1
	n2 := 2
	n3 := 3
	jeruk := "Rp. 15.000"
	hargaJeruk := 15000
	fmt.Println("Harga Jeruk ",jeruk)
	fmt.Println("--------------------")

	identitas := "alucard_1202190353_si4403"

	upperString := strings.ToUpper(identitas)

	fmt.Println(upperString)
	fmt.Println("--------------------")

	///////////////////////////////////////////

	name := "Haris"
	fmt.Println("%s buy orange with price : %d", name, jeruk)
	fmt.Printf("%s buy orange with price : Rp.%d \n", name, hargaJeruk)
	word := fmt.Sprintf("(this is sprintf) %s buy orange with price : Rp.%d\n", name, hargaJeruk)
	fmt.Print(word)

	/////////////////////////////////////////////
	res2 := fmt.Sprintf("There are %[2]d oranges %d apples %[1]d plums\n", n1, n2, n3)
    fmt.Println(res2)

	/////////////////////////////////////////////
	fmt.Printf("%10f\n", 16.540) 
	/* The width is the minimum number of runes to be output. 
	If the width is greater than the value, it is padded with spaces.
	%10 -> width = 10
	*/


	fmt.Printf("%0.f\n", 16.540)
    fmt.Printf("%0.2f\n", 16.540)
    fmt.Printf("%.3f\n", 16.540)
    fmt.Printf("%.5f\n", 16.540)
}