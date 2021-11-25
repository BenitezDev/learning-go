package main

import (
	"flag"
	"fmt"
	"github.com/BenitezDev/learning-go/quiz-game/problems"
)

func main() {

	helpPtr := flag.Bool("h", false, "display help options")
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30)")

	flag.Parse()

	fmt.Println("help:", *helpPtr)
	fmt.Println("csv:", *csvPtr)
	fmt.Println("limit:", *limitPtr)

	if *helpPtr {
		fmt.Println("Usage of the quiz game:")
		fmt.Println("\t-csv string")
		fmt.Println("\t\ta csv file in the format of 'question,answer' (default \"problems.csv\")")
		fmt.Println("\t-limit int")
		fmt.Println("\t\tthe time limit for the quiz in seconds (default 30)")
		return
	}

	qa, _ := problems.LoadProblemsFromCSV(*csvPtr)
	fmt.Println(qa)

}
