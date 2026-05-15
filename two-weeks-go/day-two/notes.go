/*
* Objectives:
* - Arrays and Slices
* - Append, len, cap
* - Maps
* - Range-based for loops
* - Variadic Functions
 */
package main

import (
	"fmt"
)

/* Variadic function */
func sum(nums ...int) (total int) {
	fmt.Println(nums)
	for _, num := range nums {
		total += num
	}
	return
}
func main() {
	/* Arrays */
	myArr := [3]string{"hello", "from", "myArr"}
	fmt.Println(myArr)

	/* Access by index */
	fmt.Println(myArr[2])

	/* Change */
	myArr[2] = "changed"
	fmt.Println(myArr[2])

	/* Initializing Arrays */
	a1 := [5]int{}
	a2 := [5]int{1}
	a3 := [5]int{1: 10, 2: 20}
	fmt.Println(a1, a2, a3)

	/* len() */
	fmt.Println("Length of myArr: ", len(myArr))

	/* Slices */
	mySlice := []int{}
	fmt.Println(mySlice)
	/* cap() */
	s2 := []int{1, 2, 4}
	fmt.Println("Capacity of my slice: ", cap(s2))
	/* Slice from make function */
	s3 := make([]string, 5, 10)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	s4 := make([]string, 5)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))
	/* append */
	s4 = append(s4, "end?", "of the road")
	fmt.Println(s4)
	s3 = append(s3, "added")

	/* append slices */
	s5 := append(s3, s4...)
	fmt.Println(s5)

	/* Maps */
	fmt.Println("MAPS: ")
	var m1 = map[string]string{"key": "value", "go": "maps!"}
	fmt.Println(m1)

	m2 := make(map[int]string)
	m2[1] = "Alice"
	m2[2] = "Bob"
	fmt.Println(m2)
	fmt.Println(m2[2])

	m2[3] = "Charlie"
	fmt.Println(m2)
	delete(m2, 3)
	fmt.Println(m2)

	/* check for elements in map*/
	charlie, ok := m2[3]
	fmt.Println(charlie, ok)

	/* iterate over map */
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println(sum(1, 2, 4, 5, 91))

}
