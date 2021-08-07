package calc

import (
	"math"
	"testing"
)

func compareExpressionAns(exp string, ans float64, t *testing.T) {
	res, err := Calc(exp)
	if err != nil {
		t.Errorf("%s should not be error", exp)
	} else if math.Abs(res-ans) > 1e-6 {
		t.Errorf("%s, Get %v, Expect: %v", exp, res, ans)
	}
}

func testWrongExpression(exp string, t *testing.T) {
	_, err := Calc(exp)
	if err == nil {
		t.Errorf("%s should be error", exp)
	}
}

func TestCalc(t *testing.T) {
	compareExpressionAns("1+2*5", 11, t)
	compareExpressionAns("(1+2)*5", 15, t)
	compareExpressionAns("(1 +2 )*3", 9, t)
	compareExpressionAns("(-1)", -1, t)
	compareExpressionAns("((-1))", -1, t)
	compareExpressionAns("((-1))*(-2)", 2, t)
	compareExpressionAns("((-3.14))*(-12)", 3.14*12, t)
	testWrongExpression("-1", t)
	testWrongExpression("1/0", t)
	testWrongExpression("-1+1", t)
}
