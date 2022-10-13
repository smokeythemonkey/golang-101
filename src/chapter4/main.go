// Variables
package main

import "fmt"

// func main(){
// 	var x string = "Hello World"
// 	fmt.Println(x)
// }

// variables are mutable in Go
// func main()  {
// 	var x string
// 	x = "first"
// 	fmt.Println(x)
// 	x = x + "second"
// 	fmt.Println(x)
// }
// vars should be in camelCase := is short had for defining a var, letting compiler infer type
// func main() {
	// doggoName := "Boomer"
	// fmt.Println("My dog's name is", doggoName)
// }

// Scope of vars

// var greetingMessage string = "Hello World" 
// func main()  {
// 	fmt.Println(greetingMessage)
// 	printSomeVariable()
// }

// func printSomeVariable(){
// 	fmt.Println(greetingMessage)
// }

// Constants
//cant be reassigned
// func main(){
// 	const greetingMessage string = "Hello World"
// 	greetingMessage = "HAHAHAH"
// 	fmt.Println(greetingMessage)
// }

// func main()  {
// 	fmt.Println("Enter a number: ")
// 	var userInput float64
// 	fmt.Scanf("%f", &userInput)

// 	computedOutput := userInput * 2

// 	fmt.Println(computedOutput)
// }

func main()  {
	const(
		FAHRENHEIT_ZERO = 32
		CONVERSION_FACTOR = 1.8
	)
	fmt.Println("Enter temperature in Fahrenheit (F):" )

	var userInputFahrenheit float32
	fmt.Scanf("%f", &userInputFahrenheit)

	temperatureCelcius := (userInputFahrenheit - FAHRENHEIT_ZERO) / CONVERSION_FACTOR
	fmt.Println("The temperature converted into Celcius is: ", temperatureCelcius)
}