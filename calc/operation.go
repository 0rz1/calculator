package calc

type operation struct {
	priority  int
	calc_call func(float64, float64) float64
}

var all_operations map[string]operation = map[string]operation{
	"+": {1, func(a, b float64) float64 { return a + b }},
	"-": {1, func(a, b float64) float64 { return a - b }},
	"*": {2, func(a, b float64) float64 { return a * b }},
	"/": {2, func(a, b float64) float64 { return a / b }},
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
