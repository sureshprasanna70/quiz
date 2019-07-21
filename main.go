package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	for i, line := range lines {

		probStruct[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return probStruct
}

func main() {
	correct := 0
	csvFilename := flag.String("csv", "ques.csv", "Location of questions in csv format")
	timeLimit := flag.Int("limit", 30, "Time limit to answer a question")
	flag.Parse()
	//parse the flags before using it
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	answerCh := make(chan string)
	dat, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("failed to open %s\n", *csvFilename)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(dat))
	lines, err := reader.ReadAll()
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("\nQuestion #%d : %s =", i+1, p.question)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d\n", correct, len(problems))

			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct = correct + 1
			}
		}

	}
	fmt.Printf("\nYou got %d out of %d\n", correct, len(problems))
}
