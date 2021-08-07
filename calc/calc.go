package calc

import "strconv"

type wordType int

const (
	wt_parentheses wordType = iota
	wt_number      wordType = iota
	wt_operation   wordType = iota
	wt_invalid     wordType = iota
)

type word struct {
	val string
	wt  wordType
}

func (w *word) Valid() bool {
	switch w.wt {
	case wt_parentheses:
		return w.val == "(" || w.val == ")"
	case wt_number:
		_, err := strconv.ParseFloat(w.val, 64)
		return err != nil
	case wt_operation:
		_, ok := pickOp(w.val)
		return ok
	default:
		return false
	}
}

func splitWord(bs []byte) (words []word) {
	length := len(bs)
	for i := 0; i < length; {
		opbs := pickOpBs(bs[i:])
		if opbs != nil {
			words = append(words, word{string(opbs), wt_operation})
			i += len(opbs)
		} else if bs[i] == '(' {
			words = append(words, word{string(bs[i : i+1]), wt_parentheses})
			i++
		} else if bs[i] == ')' {
			wordsLen := len(words)
			if wordsLen >= 3 &&
				words[wordsLen-1].wt == wt_number &&
				words[wordsLen-2].val == "-" &&
				words[wordsLen-3].val == "(" {
				words[wordsLen-2].wt = wt_number
				words[wordsLen-2].val += words[wordsLen-1].val
				words[wordsLen-1] = word{string(bs[i : i+1]), wt_parentheses}
			} else {
				words = append(words, word{string(bs[i : i+1]), wt_parentheses})
			}
			i++
		} else if bs[i] >= '0' && bs[i] <= '9' {
			j := i + 1
			for j < length && (bs[j] >= '0' && bs[j] <= '9' || bs[j] == '.') {
				j++
			}
			words = append(words, word{string(bs[i:j]), wt_number})
			i = j
		} else if bs[i] == ' ' {
			i++
		} else {
			words = append(words, word{string(bs[i : i+1]), wt_invalid})
			i++
		}
	}
	return
}

func checkGrammer(words []word) bool {
	pCnt := 0
	readNumOrOp := true
	for _, w := range words {
		if !w.Valid() {
			return false
		}
		switch w.wt {
		case wt_number:
			if !readNumOrOp {
				return false
			}
			readNumOrOp = false
		case wt_operation:
			if readNumOrOp {
				return false
			}
			readNumOrOp = true
		case wt_parentheses:
			shouldReadNumOrOp := w.val == "("
			if shouldReadNumOrOp != readNumOrOp {
				return false
			}
			if w.val == "(" {
				pCnt++
			} else {
				pCnt--
				if pCnt < 0 {
					return false
				}
			}
		default:
			return false
		}
	}
	return !readNumOrOp && pCnt == 0
}

func makeTree(words []word) node {
	priority := 0
	var tree node
	for _, w := range words {
		switch w.wt {
		case wt_parentheses:
			if w.val[0] == '(' {
				priority += 100
			} else {
				priority -= 100
			}
		case wt_number:
			val, err := strconv.ParseFloat(w.val, 64)
			if err != nil {
				panic("wt_number err")
			}
			nn := &numNode{val}
			if tree == nil {
				tree = nn
			} else {
				tree.Concat(nn)
			}
		case wt_operation:
			op, ok := pickOp(w.val)
			if !ok {
				panic("wt_operation err")
			}
			op.priority += priority
			opn := &opNode{op, nil, nil}
			tree.Concat(opn)
		}
	}
	return tree
}

func Calc(expression string) float64 {
	bs := []byte(expression)
	words := splitWord(bs)
	if checkGrammer(words) {
		tree := makeTree(words)
		if tree != nil {
			return tree.Calc()
		}
	}
	return 0
}
