package main

import "fmt"

// arrays cannot change lengths, values are mutable
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

// better, average calculation is dynamic
// func main() {
// 	x := [5]float64{98, 93, 77, 82, 83}
// 	var sum float64 = 0
// 	for i := 0; i < len(x); i++ {
// 		sum += x[i]
// 	}
// 	fmt.Println(sum / float64(len(x)))
// }

// using the special loop form to simplify it
// range tells loop which object to iterate over
// value == x[i]
// func main() {
// x := [5]float64{98, 93, 77, 82, 83}
// var total float64 = 0
// for _, value := range x {
// total += value
// }
// fmt.Println(total / float64(len(x)))
// }

//SLICES append

// func main()  {
// 	slice1 := []int{1,2,3}
// 	slice2 := append(slice1, 4, 5)
// 	fmt.Println(slice1,slice2)
// }

// copy
// slice2 is only 2 elements big, hence only first 2 elements of slice1 are copied
// func main()  {
// 	slice1 := []int{1,2,3}
// 	slice2 := make([]int,2)
// 	copy(slice2,slice1)
// 	fmt.Println(slice1,slice2)
// }

//MAPS == python dict
// gives you a runtime error! need to make (init) the map first
// func main()  {
// 	var x map[string]int
// 	x["key1"] = 10
// 	fmt.Println(x)
// }

// func main()  {
// 	x := make(map[string]int)
// 	x["key1"] = 10
// 	fmt.Println(x)

// 	y := make(map[int]int)
// 	y[2] = 2
// 	fmt.Println(y)

// 	y[34] = 34
// 	fmt.Println(y)
// 	delete(y,2)
// 	fmt.Println(y)
// }

// func main() {
// 	periodic_elements := make(map[string]string)
// 	periodic_elements["H"] = "Hydrogen"
// 	periodic_elements["He"] = "Helium"
// 	periodic_elements["Li"] = "Lithium"
// 	periodic_elements["Be"] = "Beryllium"
// 	periodic_elements["B"] = "Boron"
// 	periodic_elements["C"] = "Carbon"
// 	periodic_elements["N"] = "Nitrogen"
// 	periodic_elements["O"] = "Oxygen"
// 	periodic_elements["F"] = "Fluorine"
// 	periodic_elements["Ne"] = "Neon"

// 	// unknown elements return 0
// 	// fmt.Println(periodic_elements["Un"])
// 	// check if element is there
// 	name, ok := periodic_elements["Un"]
// 	fmt.Println(name, ok)

// 	if name, ok := periodic_elements["Un"]; ok {
// 		fmt.Println(name, ok)
// 	}
// }

// MAPS IN MAPS

func main() {
	periodic_elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := periodic_elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
