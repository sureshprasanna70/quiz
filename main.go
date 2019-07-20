package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	tLimit   int
}

func parseLines(lines [][]string) []problem {
	probStruct := make([]problem, len(lines))
	for i, line := range lines {
		intTime, _ := strconv.Atoi(line[2])
		probStruct[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
			tLimit:   intTime,
		}
	}
	return probStruct
}

func main() {
	correct := 0
	csvFilename := flag.String("csv", "ques.csv", "Location of questions in csv format")
	//timeLimit := flag.Int("limit", 30, "Time limit to answer a question")
	flag.Parse()
	//parse the flags before using it
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
		fmt.Printf("\nYou have got %d seconds to answer this question", p.tLimit)
		timer := time.NewTimer(time.Duration(p.tLimit) * time.Second)
		fmt.Printf("\nQuestion #%d : %s =", i+1, p.question)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			continue
		case answer := <-answerCh:
			if answer == p.answer {
				correct = correct + 1
			}
		}

	}
	fmt.Printf("You got %d out of %d\n", correct, len(problems))
}
