package main

import "fmt"




// like a python list
// func main()  {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x)
// }

// loop through an array, the dumb way
// func main()  {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 83
// 	x[2] = 12
// 	x[3] = 34
// 	x[4] = 15

// 	var sum float64 = 0
// 	for i:=0;i<5; i++{
// 		sum += x[i]
// 	}
// 	fmt.Println(sum / 5)
// }

func main()  {
	x := [5]float64{98,83, 12,34,15}
	var sum float64 = 0
	for i :=0; i<len(x); i++{
		sum += x[i]
	}
	fmt.Println(sum / float64(len(x)))
}

