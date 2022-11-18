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

	fmt.Println(x == z, z > y)

	var massiv1 = [15]int{0, 1, 2, 3}
	var massiv2 = [5]int{0, 1, 2, 3, 4}
	massiv3 := [...]int{0, 1, 2}

	fmt.Println(massiv1, "\n", massiv2, "\n", massiv3, "\nmassiv1[0] == massiv3[0]: ", massiv1[0] == massiv3[0])

	massiv4 := [...]int{3: 1, 4: 2, 5: 3}
	fmt.Println(massiv4)

	if x > x {
		println("IF IF IF IF\nIF IF IF IF")
	} else if x == x {
		println("ELSEIF ELSEIF ELSEIF\nELSEIF ELSEIF ELSEIF")
	} else {
		println("ELSE")
	}

	a := 8
	switch a {
	case 0:
		println("it's zero")
	case 1:
		println("it's one")
	case 2:
		println("it's two")
	case 3, 4, 5:
		println("it's three, four or five")
	default:
		println("unknown number")
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	names := [4]string{"Alice", "Paul", "Tom", "Jack"}
	fmt.Println(names)

	for _, value := range names {
		fmt.Println(value)
	}
	fmt.Println("names: ", len(names))
}
