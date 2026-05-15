/*
* Date: May 15th, 2026
* Author:
* Description: Read student names + scores from stdin. Store in a map.
* Print the class average, highest scorer, and a sorted grade list.
 */
package main

import (
	"fmt"
	"strings"
)

func calculateClassAverage(scores map[string]int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

func main() {
	fmt.Println("=== Student Stats ===")

	var userDecision string
	studentScores := make(map[string]int)

	for {
		if strings.ToLower(userDecision) == "n" {
			break
		}
		var name string
		var score int
		fmt.Println("Enter Student Name: ")
		fmt.Scan(&name)
		fmt.Println("Enter Student Score: ")
		fmt.Scan(&score)
		studentScores[name] = score
		fmt.Println("Enter another student? (Y/n):")
		fmt.Scan(&userDecision)
	}

	fmt.Println(studentScores)
	avg := calculateClassAverage(studentScores)
	fmt.Printf("Class Average:\t %f", avg)
}
