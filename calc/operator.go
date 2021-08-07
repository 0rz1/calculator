package calc

type operator struct {
	priority  int
	calc_call func(float64, float64) (float64, error)
}

var all_operators map[string]operator = map[string]operator{
	"+": {1, func(a, b float64) (float64, error) { return a + b, nil }},
	"-": {1, func(a, b float64) (float64, error) { return a - b, nil }},
	"*": {2, func(a, b float64) (float64, error) { return a * b, nil }},
	"/": {2, func(a, b float64) (float64, error) {
		if b == 0 {
			err := CalcErr("divide zero")
			return 0, &err
		}
		return a / b, nil
	}},
}

func pickOp(opStr string) (operator, bool) {
	for k, op := range all_operators {
		if k == opStr {
			return op, true
		}
	}
	return operator{}, false
}

func pickOpBsPre(bs []byte) []byte {
	for k, _ := range all_operators {
		ks := []byte(k)
		if len(ks) <= len(bs) && string(bs[:len(ks)]) == k {
			return ks
		}
	}
	return nil
}
