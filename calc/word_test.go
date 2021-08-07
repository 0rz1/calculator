package calc

import (
	"testing"
)

func TestWordValid(t *testing.T) {
	for _, s := range []string{"(", ")"} {
		w := word{s, wt_parentheses}
		if !w.Valid() {
			t.Errorf("%v is valid", s)
		}
	}
	for _, s := range []string{"a", "((", "))"} {
		w := word{s, wt_parentheses}
		if w.Valid() {
			t.Errorf("%v is not valid", s)
		}
	}
	for _, s := range []string{"1", "1.0", "1.", "-1", "-123.456"} {
		w := word{s, wt_number}
		if !w.Valid() {
			t.Errorf("%v is valid", s)
		}
	}
	for _, s := range []string{"(1)", "1a", "1)", "1-1", "2.-1", "1.1.2"} {
		w := word{s, wt_number}
		if w.Valid() {
			t.Errorf("%v is not valid", s)
		}
	}
	for _, s := range []string{"+", "-", "*", "/"} {
		w := word{s, wt_operator}
		if !w.Valid() {
			t.Errorf("%v is valid", s)
		}
	}
	for _, s := range []string{"*/", "--", "++", "a"} {
		w := word{s, wt_operator}
		if w.Valid() {
			t.Errorf("%v is not valid", s)
		}
	}
	for _, s := range []string{"1", "+", ")", "ab"} {
		w := word{s, wt_invalid}
		if w.Valid() {
			t.Errorf("%v is not valid when wt_invalid", s)
		}
	}
}

func compareSplitWords(expression string, expects []string, t *testing.T) {
	words := splitWord([]byte(expression))
	if len(words) != len(expects) {
		t.Fatalf("12+-1)1.1 should split to %d but %d", len(expects), len(words))
	}
	for i, w := range words {
		if w.val != expects[i] {
			t.Errorf("%v", words[0].val)
		}
	}
}

func TestSplitWord(t *testing.T) {
	compareSplitWords("12+-1)1.1", []string{"12", "+", "-", "1", ")", "1.1"}, t)
	compareSplitWords("1 2+-1)1. 1", []string{"1", "2", "+", "-", "1", ")", "1.", "1"}, t)
	compareSplitWords("-1", []string{"-", "1"}, t)
	compareSplitWords("(-1)", []string{"(", "-1", ")"}, t)
}

func TestCheckGrammer(t *testing.T) {
	words := []word{
		{"(", wt_parentheses},
		{"(", wt_parentheses},
		{")", wt_parentheses},
		{")", wt_parentheses},
	}
	if err := checkGrammer(words); err == nil {
		t.Errorf("(()) invalid ")
	}
	words = []word{
		{"(", wt_parentheses},
		{"(", wt_parentheses},
		{"1", wt_number},
		{")", wt_parentheses},
		{")", wt_parentheses},
	}
	if err := checkGrammer(words); err != nil {
		t.Errorf("((1)) valid %v", err)
	}
	words = []word{
		{"(", wt_parentheses},
		{"(", wt_parentheses},
		{"1", wt_number},
		{")", wt_parentheses},
		{")", wt_parentheses},
		{"(", wt_parentheses},
	}
	if err := checkGrammer(words); err == nil {
		t.Errorf("((1))( invalid")
	}
	words = []word{
		{"5", wt_number},
		{"-", wt_operator},
		{"5", wt_number},
	}
	if err := checkGrammer(words); err != nil {
		t.Errorf("5-5 valid %v", err)
	}
	words = []word{
		{"5", wt_number},
		{"(", wt_parentheses},
		{"-", wt_operator},
		{")", wt_parentheses},
		{"5", wt_number},
	}
	if err := checkGrammer(words); err == nil {
		t.Errorf("5(-)5 invalid")
	}
}

func TestMakeTree(t *testing.T) {
	ws := splitWord([]byte("2+3*4"))
	tree := makeTree(ws)
	if tree == nil {
		t.Fatalf("2+3*4, tree is not nil")
	} else {
		res, err := tree.Calc()
		if err == nil {
			t.Logf("2+3*4 = %v", res)
		}
	}
	tn, ok := tree.(*opNode)
	if !ok {
		t.Fatal("(2+3)*4 tree should be opNode")
	} else if tn.op.priority != 1 {
		t.Errorf("* Get %v, Expect %d", tn.op.priority, 1)
	}
	ws = splitWord([]byte("(2+3)*4"))
	tree = makeTree(ws)
	if tree == nil {
		t.Fatal("(2+3)*4 tree is not nil")
	}
	tn, ok = tree.(*opNode)
	if !ok {
		t.Fatal("(2+3)*4 tree should be opNode")
	} else if tn.op.priority != 2 {
		t.Errorf("* Get %v, Expect %d", tn.op.priority, 2)
	}
}
