package calc

import "fmt"

type CalcErr string

func (e *CalcErr) Error() string {
	return fmt.Sprintf("Calc error: %v", *e)
}

type WordValidErr word

func (e *WordValidErr) Error() string {
	switch e.wt {
	case wt_number:
		return fmt.Sprintf("number %v is not valid", e.val)
	case wt_operation:
		return fmt.Sprintf("op %v is not valid", e.val)
	case wt_parentheses:
		return fmt.Sprintf("parentheses %v is not valid", e.val)
	default:
		return fmt.Sprintf("%v shoud not exist", e.val)
	}
}

type GrammerErr string

func (e *GrammerErr) Error() string {
	return fmt.Sprintf("Grammer error: %v", *e)
}
