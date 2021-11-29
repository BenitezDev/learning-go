package main

import (
	"flag"
	"fmt"
	. "github.com/BenitezDev/learning-go/quiz-game/problems"
	"strings"
	"time"
)

func printHelp() {
	fmt.Println("Usage of the quiz game:")
	fmt.Println("\t-csv string")
	fmt.Println("\t\ta csv file in the format of 'question,answer' (default \"problems.csv\")")
	fmt.Println("\t-limit int")
	fmt.Println("\t\tthe time limit for the quiz in seconds (default 30)")
}

func EndGame(hits, totalQuestions *int) {

	fmt.Println()
	fmt.Printf("You scored %d out of %d.\n", *hits, *totalQuestions)

	// wait for user to press a key to finish the quiz
	_, _ = fmt.Scanln()
}

func AskUser(p *Problem, index *int, ch chan bool) {

	fmt.Printf("Problem #%d: %s=", *index+1, p.Question)
	var input string
	_, _ = fmt.Scanln(&input)

	if strings.Compare(p.Answer, input) == 0 {
		ch <- true
	}
	ch <- false
}

func main() {

	helpPtr := flag.Bool("h", false, "display help options")
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30)")

	flag.Parse()

	if *helpPtr {
		printHelp()
		return
	}

	qa, err := LoadProblemsFromCSV(*csvPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	hitChannel := make(chan bool)

	hits := 0
	totalQuestions := len(qa)

	for i := 0; i < totalQuestions; i++ {

		go AskUser(&qa[i], &i, hitChannel)

		select {

		case hit := <-hitChannel:
			if hit {
				hits++
			}

		case <-time.After(time.Duration(*limitPtr) * time.Second):
			EndGame(&hits, &totalQuestions)
			return
		}
	}

}
