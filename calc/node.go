package calc

type numNode struct {
	val float64
}

func (n numNode) Calc() (float64, error) {
	return n.val, nil
}

func (n numNode) Priority() int {
	return 1 << 30
}

func (n *numNode) Merge(x int, ch node) {
	panic("Merge Err: numNode should not merge child")
}

type opNode struct {
	op    operator
	left  node
	right node
}

func (n opNode) Calc() (float64, error) {
	lv, le := n.left.Calc()
	if le != nil {
		return 0, le
	}
	rv, re := n.right.Calc()
	if re != nil {
		return 0, re
	}
	return n.op.calc_call(lv, rv)
}

func (n opNode) Priority() int {
	return n.op.priority
}

func (n *opNode) Merge(x int, ch node) {
	if x == 0 {
		n.left = concat(ch, n.left)
	} else if x == 1 {
		n.right = concat(n.right, ch)
	} else {
		panic("Merge Err: x out of range")
	}
}

type node interface {
	Calc() (float64, error)
	Priority() int
	Merge(int, node)
}

func concat(a node, b node) node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Priority() < b.Priority() {
		a.Merge(1, b)
		return a
	} else {
		b.Merge(0, a)
		return b
	}
}
