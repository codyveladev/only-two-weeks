/*
* Date: May 15th, 2026
* Author:
* Description: Read student names + scores from stdin. Store in a map.
* Print the class average, highest scorer, and a sorted grade list.
 */
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFromStdin() map[string]int {

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
	return studentScores
}

func loadStudentsFromFile(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	studentScores := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Printf("skipping invalid line: %s\n", line)
			continue
		}
		name := strings.TrimSpace(parts[0])
		score, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			fmt.Printf("skipping invalid score for %s\n", name)
			continue
		}
		studentScores[name] = score
	}

	return studentScores, scanner.Err()
}

func calculateClassAverage(scores map[string]int) float64 {
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

func findTopScorers(scores map[string]int) map[string]int {
	highScore := math.MinInt
	result := make(map[string]int)
	for name, score := range scores {
		if score > highScore {
			result = make(map[string]int)
			result[name] = score
			highScore = score
		} else if score == highScore {
			result[name] = score
		}
	}
	return result
}

func sortScores(scores map[string]int) []string {
	names := []string{}
	for name := range scores {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		if scores[names[i]] == scores[names[j]] {
			return names[i] < names[j] // alphabetical as tiebreaker
		}
		return scores[names[i]] > scores[names[j]]
	})

	return names
}

func main() {
	fmt.Println("=== Student Stats ===")
	studentScores := make(map[string]int)

	var err error

	if len(os.Args) > 1 {
		studentScores, err = loadStudentsFromFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		studentScores = readFromStdin()
	}

	fmt.Println("== Analysis ==")
	avg := calculateClassAverage(studentScores)
	fmt.Printf("Class Average:\t %.2f\n", avg)

	highScores := findTopScorers(studentScores)
	fmt.Println("== Top Scorers ==")
	for name, score := range highScores {
		fmt.Printf("%s:\t %d\n", name, score)
	}
	fmt.Println("== Leaderboard ==")
	sortedNames := sortScores(studentScores)
	for _, name := range sortedNames {
		fmt.Printf("%s:\t %d\n", name, studentScores[name])
	}
}
