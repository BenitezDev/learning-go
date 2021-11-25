package main

import (
	"flag"
	"fmt"
	"github.com/BenitezDev/learning-go/quiz-game/problems"
	"strings"
)

func printHelp() {
	fmt.Println("Usage of the quiz game:")
	fmt.Println("\t-csv string")
	fmt.Println("\t\ta csv file in the format of 'question,answer' (default \"problems.csv\")")
	fmt.Println("\t-limit int")
	fmt.Println("\t\tthe time limit for the quiz in seconds (default 30)")
}

func main() {

	helpPtr := flag.Bool("h", false, "display help options")
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")

	flag.Parse()

	if *helpPtr {
		printHelp()
		return
	}

	qa, err := problems.LoadProblemsFromCSV(*csvPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalQuestions := len(qa)

	hits := 0
	var input string

	for i := 0; i < totalQuestions; i++ {

		fmt.Printf("Problem #%d: %s=", i+1, qa[i].Question)

		_, _ = fmt.Scanln(&input)

		if strings.Compare(qa[i].Answer, input) == 0 {
			hits++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", hits, totalQuestions)
}
