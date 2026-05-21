package utils

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/codyeladev/day-five-project/models"
	"github.com/fatih/color"
)

// register all helpers
func readStudentsFromStdin() []models.Student {

	var userDecision string
	studentScores := []models.Student{}
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
		student := models.Student{Name: name, Score: score}
		studentScores = append(studentScores, student)
		fmt.Println("Enter another student? (Y/n):")
		fmt.Scan(&userDecision)
	}
	return studentScores
}

func calculateClassAverage(students []models.Student) float64 {
	total := 0
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func findTopScorers(students []models.Student) []models.Student {
	highScore := math.MinInt
	result := []models.Student{}
	for _, student := range students {
		if student.Score > highScore {
			result = []models.Student{student}
			highScore = student.Score
		} else if student.Score == highScore {
			result = append(result, student)
		}
	}
	return result
}

func printTopScorers(students []models.Student) {
	for _, student := range students {
		fmt.Printf("%s:\t%d\n", student.Name, student.Score)
	}
}

func sortScores(students []models.Student) []models.Student {
	sort.Slice(students, func(i, j int) bool {
		if students[i].Score == students[j].Score {
			return students[i].Name < students[j].Name
		}
		return students[i].Score > students[j].Score
	})
	return students
}

func printScoresLeaderboard(students []models.Student) {
	for i, student := range students {
		fmt.Printf("%d. %s - %d\n", i+1, student.Name, student.Score)
	}
}

func RunStudentScoreAnalysis() {
	color.Blue("== Welcome to Student Analysis Tool ==")
	students := readStudentsFromStdin()
	color.Blue("Loading Stats...")
	color.Green("Class Average: %.2f", calculateClassAverage(students))
	topScores := findTopScorers(students)
	color.Blue("Top Scores:")
	printTopScorers(topScores)
	color.Blue("Score Leader Board:")
	printScoresLeaderboard(sortScores(students))
}
