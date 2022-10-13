package main

import "fmt"
// simple for loop go only has one for loop that can be used in diff ways
// func main()  {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i += 1
// 	}
// }

// like a python for loop and enumerate
// func main()  {
// 	for i := 1; i <= 10 ; i++{
// 		fmt.Println(i)
// 	}
// }

// if statements, syntax of the else block seems weird
// func main()  {
// 	for i := 1; i <= 10; i++ {
// 		if i % 2 == 0 {
// 			fmt.Println("even")
// 		} else {
// 			fmt.Println("odd")
// 		}
// 	}
// }

// switch

func main()  {
	i := 0
	switch i {
	case 0: fmt.Println("Case Zero")
	case 1: fmt.Println("Case One")
	default: fmt.Println("Unknown Number")
	}
}