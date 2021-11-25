package problems

import (
	"io/fs"
)

type Problem struct {
	Question string
	Answer   string
}

type Problems []Problem

func LoadProblemsFromCSV(fileSystem fs.FS) Problems {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var problems Problems
	for range dir {
		problems = append(problems, Problem{})
	}
	return problems
}
