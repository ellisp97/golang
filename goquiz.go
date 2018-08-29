package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "CSV File: Question , Answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i,problem := range problems {
		fmt.Printf("Problem #%d: %s = " ,i+1 , problem.q)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out %d.\n", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == problem.a{
				correct++
			}
		}
	}
}


func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(errmsg string) {
	fmt.Println(errmsg)
	os.Exit(1)
}