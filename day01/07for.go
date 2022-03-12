package main

import "fmt"


func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 无限循环
	// for {

	// }

	// for range
	s := "Hello姜昌雄"
	for i, v := range s {
		fmt.Println(i, v)
	}


	// 打印九九乘法表

	for i := 1; i <= 9; i++ {
		for j := 1; j <=i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i * j)
		}
		fmt.Println()
	}

	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i * j)
		}
		fmt.Println()
	}
}