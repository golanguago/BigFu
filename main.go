package main

import "fmt"

// Std_func
func main() {
	fmt.Print("string")
	fmt.Println("_s")

	var stroka1 string
	var x, y int
	x = 5
	y = 12

	stroka1 = "text of stroka1"
	fmt.Println(stroka1, "\t", x, "\t", y)

	var stroka2 = "text of stroka2"
	fmt.Println(stroka2)

	z := 67
	stroka3 := "text of stroka3"
	fmt.Println(stroka3, "\t", z)

	xyz := x + y + z
	fmt.Println(xyz)
}
