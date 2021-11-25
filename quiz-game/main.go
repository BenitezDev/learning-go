package main

import (
	"flag"
	"fmt"
	"github.com/BenitezDev/learning-go/quiz-game/problems"
	"strings"
)

func main() {

	helpPtr := flag.Bool("h", false, "display help options")
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")

	flag.Parse()

	if *helpPtr {
		fmt.Println("Usage of the quiz game:")
		fmt.Println("\t-csv string")
		fmt.Println("\t\ta csv file in the format of 'question,answer' (default \"problems.csv\")")
		fmt.Println("\t-limit int")
		fmt.Println("\t\tthe time limit for the quiz in seconds (default 30)")
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

		fmt.Print(qa[i].Question, "=")

		_, _ = fmt.Scanln(&input)

		if strings.Compare(qa[i].Answer, input) == 0 {
			hits++
			fmt.Print("  Correct!")
		} else {
			fmt.Print("  Incorrect!")
		}

		fmt.Println(" (", hits, "/", totalQuestions, ") hits")
		input = ""
	}

}
