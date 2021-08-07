package calc

type operation struct {
	priority  int
	calc_call func(float64, float64) (float64, error)
}

var all_operations map[string]operation = map[string]operation{
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

func pickOp(opStr string) (operation, bool) {
	for k, op := range all_operations {
		if k == opStr {
			return op, true
		}
	}
	return operation{}, false
}

func pickOpBs(bs []byte) []byte {
	for k, _ := range all_operations {
		ks := []byte(k)
		if len(ks) <= len(bs) && string(bs[:len(ks)]) == k {
			return ks
		}
	}
	return nil
}
