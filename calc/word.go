package calc

import (
	"fmt"
	"strconv"
)

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
		return err == nil
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

func checkGrammer(words []word) error {
	pCnt := 0
	readNumOrOp := true
	for _, w := range words {
		if !w.Valid() {
			err := WordValidErr(w)
			return &err
		}
		switch w.wt {
		case wt_number:
			if !readNumOrOp {
				err := GrammerErr("need op but number")
				return &err
			}
			readNumOrOp = false
		case wt_operation:
			if readNumOrOp {
				err := GrammerErr("need number but op")
				return &err
			}
			readNumOrOp = true
		case wt_parentheses:
			shouldReadNumOrOp := w.val == "("
			if shouldReadNumOrOp != readNumOrOp {
				err := GrammerErr(fmt.Sprintf("before %v is not right type", w.val))
				return &err
			}
			if w.val == "(" {
				pCnt++
			} else {
				pCnt--
				if pCnt < 0 {
					err := GrammerErr(") is too much")
					return &err
				}
			}
		default:
			err := GrammerErr(fmt.Sprintf("w.val is not valid type", w.val))
			return &err
		}
	}
	if pCnt != 0 {
		err := GrammerErr(") is less than (")
		return &err
	} else if readNumOrOp {
		err := GrammerErr("last word should be number")
		return &err
	} else {
		return nil
	}
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
			tree = concat(tree, nn)
		case wt_operation:
			op, ok := pickOp(w.val)
			if !ok {
				panic("wt_operation err")
			}
			op.priority += priority
			opn := &opNode{op, nil, nil}
			tree = concat(tree, opn)
		}
	}
	return tree
}
