# Day Five Project — Student Score Analysis

A small Go CLI that collects student names and scores from stdin, then prints class stats: average, top scorer(s), and a sorted leaderboard. Output is colorized with `fatih/color`.

## Structure

```
day-five-project/
├── main.go
├── go.mod
├── go.sum
├── models/
│   └── student.go
└── utils/
    └── helpers.go
```

## Files

### [main.go](main.go)
Program entry point. The `main` package imports `utils` and calls `utils.RunStudentScoreAnalysis()` — all real work lives in the utils package.

### [go.mod](go.mod) / [go.sum](go.sum)
Go module definition (`github.com/codyeladev/day-five-project`, Go 1.26.3) and dependency checksums. Depends on `github.com/fatih/color` (with its transitive deps `mattn/go-colorable`, `mattn/go-isatty`, and `golang.org/x/sys`) for terminal colors.

### [models/student.go](models/student.go)
Defines the `Student` struct used throughout the program:
- `Name string` — student's name
- `Score int` — student's score
- `id int` — unused/unexported id field

### [utils/helpers.go](utils/helpers.go)
Holds all the program logic. Functions:

- **`readStudentsFromStdin()`** — Prompts the user in a loop for a name and score, appending each entry to a slice. Stops when the user answers `n` to "Enter another student?".
- **`calculateClassAverage(students)`** — Sums all scores and returns the mean as a `float64`.
- **`findTopScorers(students)`** — Single-pass scan that returns every student tied for the highest score.
- **`printTopScorers(students)`** — Prints the top scorers as `Name<TAB>Score`.
- **`sortScores(students)`** — Sorts students by score descending; ties broken alphabetically by name. Sorts in place and also returns the slice.
- **`printScoresLeaderboard(students)`** — Prints the sorted list as a numbered leaderboard (`1. Name - Score`).
- **`RunStudentScoreAnalysis()`** — Public entry function called by `main`. Orchestrates the flow: greet → read input → print average (green) → print top scorers → print leaderboard, with blue section headers via `fatih/color`.

## Run

```bash
cd two-weeks-go/day-five-project
go run .
```

## Improvements
I could make the packages into meaningful names and proper structure for example 

- `/display`
- `/analysis`
- `/reader`

which would break up the logic properly but this day was to more so see how packages and modules work in a small codebase. 

