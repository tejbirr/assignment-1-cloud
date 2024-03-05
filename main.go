// GOlang is CASEsensitive (C=C and not c)
package main

import "fmt" //fmt = format (u r using)

func add(x int, y int) int {
	var out = x + y
	return out
}

func add2(x, y int) (int, int) { // Returning TWO value at the same time
	var out1 = x + y
	var out2 = x - y
	return out1, out2

}

func main() { //defining afunction (u r creating)
	fmt.Println("hello bro !!")
	fmt.Print(2+5, "\n")

	// Variable and constant
	// PERMANENT = DATABASE
	// TEMPORARY = Variables   (change the value)
	var num1 int
	var num2 int
	var num3, num4 int

	num1 = 10
	num2 = 3
	num3, num4 = 2, 2
	var result = num1 + num2 + num3 + num4
	fmt.Print(result, "\n")

	result1 := 9 // var result = 9
	fmt.Print(result1, "\n")

	const num5 = 9 //CONST
	fmt.Print(num5, "\n")

	//Functions

	num6 := 5
	num7 := 5 // SAME as var num7 = 5

	result2 := add(num6, num7)
	fmt.Println(result2)

	num8 := 7
	num9 := 6

	result3, result4 := add2(num8, num9) //RETURNING TWO AT THE SMAE TIME
	fmt.Println(result3, result4)

	// LOOPS
	i := 1

	for i <= 5 {
		fmt.Println("loop", i)
		i++
	}

	// OR
	for y := 1; y <= 5; y++ {
		fmt.Println("second loop", y)
	}

	// printing 1 to 100
	for x := 1; x <= 100; x++ {
		fmt.Println(x)
	}

	// Exported Names
	// Do ONLY use Capital letters when using theirs
	// library or when creating a func outside main()

}
