package calc

import "testing"

func TestConcatNil(t *testing.T) {
	op, _ := pickOp("+")
	var numN node = &numNode{1}
	var opN node = &opNode{op, nil, nil}
	if concat(numN, nil) != numN {
		t.Log("should not be nil")
		t.Fail()
	}
	if concat(opN, nil) != opN {
		t.Log("should not be nil")
		t.Fail()
	}
	if concat(nil, numN) != numN {
		t.Log("should not be nil")
		t.Fail()
	}
	if concat(nil, opN) != opN {
		t.Log("should not be nil")
		t.Fail()
	}
	if concat(nil, nil) != nil {
		t.Log("should be nil")
		t.Fail()
	}
}

func TestConcat(t *testing.T) {
	op, _ := pickOp("+")
	var numN0 node = &numNode{1}
	var numN1 node = &numNode{2}
	var opN node = &opNode{op, nil, nil}
	opN = concat(numN0, opN)
	switch opN.(type) {
	case *numNode:
		t.Log("root should be op")
		t.Fail()
	}
	opN = concat(opN, numN1)
	switch opN.(type) {
	case *numNode:
		t.Log("root should be op")
		t.Fail()
	}
}

func TestNodeCalc(t *testing.T) {
	op, _ := pickOp("+")
	var opN node = &opNode{op, &numNode{1}, &numNode{2}}
	if ans, _ := opN.Calc(); ans != 3 {
		t.Logf("ans is %v, except 3", ans)
		t.Fail()
	}
	op1, _ := pickOp("/")
	opN = &opNode{op1, &numNode{1}, &numNode{0}}
	if _, err := opN.Calc(); err == nil {
		t.Log("divide 0, error shouldn't be nil")
		t.Fail()
	}
	var opT node = &opNode{op, &numNode{1}, &numNode{4}}
	opN = &opNode{op1, &numNode{10}, opT}
	if ans, _ := opN.Calc(); ans != 2 {
		t.Logf("ans is %v, except 2", ans)
		t.Fail()
	}
}
