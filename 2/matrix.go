package main

import "fmt"

func main() {

	// matrix := make([][]int, 10)

	// for i := 0; i < 10; i++{
	// 	matrix[i] = make([]int, 10)
	// 	for j := 0; j < 10; j++{

	// 		if i == j {
	// 		matrix[i][j] = i
	// 		}
	// 	}
	// 	fmt.Println(matrix[i])
	// }

/*  завершение цикла с помощью break, можно также for true {}
	var x int
	for {
		x++
		fmt.Println(x)
		if x > 50 {
			break
		}
	}
	*/
	rangeMatrix := []string {
		"zero",
		"first",
	}

	for index, value := range rangeMatrix {
		println(index, value)
		print(index, " ", value, "\n")
		fmt.Println(index, value)

	}


}