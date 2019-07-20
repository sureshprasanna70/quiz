package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	probStruct := make([]problem, len(lines))
	for i := range lines {
		probStruct[i] = problem{
			question: lines[i][0],
			answer:   strings.TrimSpace(lines[i][1]),
		}
	}
	return probStruct
}

func main() {
	correct := 0
	csvFilename := flag.String("csv", "ques.csv", "Location of questions in csv format")
	flag.Parse()
	dat, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("failed to open %s\n", *csvFilename)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(dat))
	lines, err := reader.ReadAll()
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Question #%d : %s \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct = correct + 1
		}
	}
	fmt.Printf("You got %d out of %d\n", correct, len(problems))
}
