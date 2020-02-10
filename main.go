//baptizing package with name
package main

import (
	"fmt"
	"strings"
)

// #2 multiplying function
// you should explicitly state types of variable: is it boolean(T/F)?, String? Integer?
// multiply(a int, b int) can be abbreviated as (a, b int)
// return value type is states as int
func multiply(a, b int) int {
	return a * b
}

// #3 uppercase function
// naked return: telling function what to return in the beginning of the code
// return variable & value is explicitly stated as (length int, uppercase string)
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I am done") // defer is "when function finishes"
	// length variable is already created on top of the function
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
	// when you use naked return, there is no need for the line below
	// return len(name), strings.ToUpper(name)
}

// #4 returning multiple strings by ...string
func repeatMe(words ...string) {
	fmt.Println(words)
}

// #5 loop 1
func superAdd(numbers ...int) int {
	// range gives you index and number
	for index, number := range numbers {
		fmt.Println(index, number) // on console, left column is index and right column is number
	}
	return 1
}

// #5 loop 2
func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 2
}

// #5 loop3
func sumTotal(numbers ...int) int {
	total := 0
	// ignoring index unlike loop 1 and loop 2
	for _, number := range numbers {
		total += number
	}
	return total
}

// #6 if
// unlike JavaScript no need for (), unlike python3, no need for :
func canIDrink(age int) bool {
	// we can make variable inside of if statement
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}

// #7 switch
// way of checking value
// replacing python if, elif, else, replacing javascript if else
func mayIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// utilizing function in go
func main() {
	// #1 hello world on console
	// var name string ="nico"
	name := "young" // same as above. := is guessing the type.
	fmt.Println("Hello" + " " + name)

	// #2 multiplying function
	fmt.Println(multiply(2, 2))

	// #3 uppercase function
	totalLength, upperName := lenAndUpper("young")
	fmt.Println(totalLength, upperName)

	// #4 returning multiple strings
	repeatMe("young", "ahn", "1", "2", "3")

	// #5 loop1
	superAdd(1, 2, 3, 4, 5, 6)

	// #5 loop2
	superAdd2(1, 2, 3, 4, 5, 6)

	// #5 loop3
	result := sumTotal(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	// #6 if statement
	fmt.Println(canIDrink(16))

	// #8 pointer
	// go allows you to do low level programming
	a := 2
	b := a // b is not going to change whehter we change variable a
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b) // print two different memory address

	c := 2
	d := &c
	c = 5
	fmt.Println(&c, d) // same memory address now
	fmt.Println(*d)    //see through memory address with *. d is also changed along with 5

	e := 2
	f := &e
	*f = 202020 // change the value of the address
	fmt.Println(e)
}

// On golang, functions are exported from different packages
// Imported function should be in upper case
