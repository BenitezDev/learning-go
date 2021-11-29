package problems

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Problem struct {
	Question string
	Answer   string
}

func (p Problem) String() string { return fmt.Sprintf("%s %s", p.Question, p.Answer) }

type Problems []Problem

func LoadProblemsFromCSV(csvPath string) (Problems, error) {

	csvFile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	var problems Problems
	for _, line := range csvLines {
		p := Problem{Question: line[0], Answer: line[1]}
		problems = append(problems, p)

	}

	return problems, nil

}
