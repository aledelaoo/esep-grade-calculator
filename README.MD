# ESEP Grade Calculator

This library is a simple utility to calculate your grades written in Go. It calculates a student's final letter grade based on assignments, exams, and essays.

## Runtime requirements:
Golang 1.24+ or newer.

## How it works
The calculator uses three types of grades:
- Assignments: 50%
- Exams: 35%
- Essays: 15%

It averages the grades for each category and uses their weights to calculate a final letter grade.

Letter grades:
- A: 90+
- B: 80-89
- C: 70-79
- D: 60-69
- F: Below 60

## Run unit tests:
```
cd grade-calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade-calculator/
go test -coverprofile=coverage.txt ./...
```

## Files
- `grade_calculator.go` - Main file for grade calculator.
- `grade_calculator_test.go` - Tests for the grade calculator.
- `go.mod` - Go module information.