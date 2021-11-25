package problems_test

import (
	"fmt"
	"github.com/BenitezDev/learning-go/quiz-game/problems"
	"reflect"
	"testing"
)

var defaultNameCSV = "problems.csv"
var defaultContentCSV = problems.Problems{
	problems.Problem{Question: "5+5", Answer: "10"},
	problems.Problem{Question: "7+3", Answer: "10"},
	problems.Problem{Question: "1+1", Answer: "2"},
	problems.Problem{Question: "8+3", Answer: "11"},
	problems.Problem{Question: "1+2", Answer: "3"},
	problems.Problem{Question: "8+6", Answer: "14"},
	problems.Problem{Question: "3+1", Answer: "4"},
	problems.Problem{Question: "1+4", Answer: "5"},
	problems.Problem{Question: "5+1", Answer: "6"},
	problems.Problem{Question: "2+3", Answer: "5"},
	problems.Problem{Question: "3+3", Answer: "6"},
	problems.Problem{Question: "2+4", Answer: "6"},
	problems.Problem{Question: "5+2", Answer: "7"},
}

func TestLoadProblemsFromCSV(t *testing.T) {
	allProblems, err := problems.LoadProblemsFromCSV(defaultNameCSV)
	fmt.Println(allProblems)

	if err != nil {
		t.Fatal(err)
	}

	got := allProblems
	want := defaultContentCSV

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
