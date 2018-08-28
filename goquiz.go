package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "CSV File: Question , Answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse CSV File")
	}
	problems := parseLines(lines)
	fmt.Println(problems)

	correct := 0
	for i,problem := range problems {
		fmt.Printf("Problem #%d: %s = \n" ,i+1 , problem.q)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == problem.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out %d", correct, len(problems))
}


func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a:line[1],
		}
	}
	return problems
}

func exit(errmsg string) {
	fmt.Println(errmsg)
	os.Exit(1)
}