/* Hello Go! */
/*
	Objectives:
		- Basic Syntax, Variables, Constants, basic types (int, float, string, bool)
		- I/O Syntax
		- if/else and loops
		- Function with multiple return values
	Project:
		- Basic CLI calculator (2 nums and 1 operator)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	//"scanner"
)

const MULTIPLIER int = 5
const URL string = "http://www.test.com"

const (
	A int = 1
	B     = 3.14
	C     = "MY CONSTANT"
)

func main() {

	/* Variable declaration */
	var a string = "Hello World"
	b := a
	var x = 2

	/* Types */
	var age int32 = 26
	var total float32 = 12.09
	var name string = "The Breeze"
	var isUploaded bool = true

	/* multiple declaration */
	c, d := 7, "6"

	fmt.Println(a, "\n", b, "\n", x)
	fmt.Println(age, total, name, isUploaded, c, d)

	/* printing constants */
	fmt.Println(MULTIPLIER, URL, A, B, C)

	/* CLI I/O*/
	var i int

	fmt.Println("Enter a number: ")
	fmt.Scan(&i)
	fmt.Println("Inputted Number: ", i)

	/* READ FROM TEXT FILE */
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	/* Read Lines */
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	/* Conditions */
	var num int

	fmt.Println("Enter a number: ")
	fmt.Scan(&num)

	if num > 5 {
		fmt.Printf("%d is greater than 5\n", num)
	} else if num < 5 {
		fmt.Printf("%d is less than 5\n", num)
	} else {
		fmt.Printf("%d is 5\n", num)
	}

	/* Nested Loops */
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	/* Range Key Word */
	fruits := [3]string{"apple", "banana", "orange"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
