package problems_test

import (
	"github.com/BenitezDev/learning-go/quiz-game/problems"
	"testing"
	"testing/fstest"
)

func TestLoadProblemsFromCSV(t *testing.T) {
	fs := fstest.MapFS{
		"problems.csv": {Data: []byte("5+5,10\n7+3,10\n1+1,2\n8+3,11\n1+2,3\n8+6,14\n3+1,4\n1+4,5\n5+1,6\n2+3,5\n3+3,6\n2+4,6\n5+2,7")},
	}
	_ = problems.Problem{Answer: "a", Question: "q"}
	allProblems := problems.LoadProblemsFromCSV(fs)

	if len(allProblems) != len(fs) {
		t.Errorf("got %d problems, wanted %d problems", len(allProblems), len(fs))
	}

}
