package calc

type numNode struct {
	val float64
}

func (n numNode) Calc() float64 {
	return n.val
}

func (n numNode) Priority() int {
	return -1
}

func (n *numNode) Concat(other node) node {
	panic("numNode UpChild should not run")
}

type opNode struct {
	op    operation
	left  node
	right node
}

func (n opNode) Calc() float64 {
	return n.op.calc_call(n.left.Calc(), n.right.Calc())
}

func (n opNode) Priority() int {
	return n.op.priority
}

func (n *opNode) Concat(other node) node {
	if other == nil {
		return n
	}
	if n == nil {
		return other
	}
	if n.Priority() >= other.Priority() {
		n.right = n.left.Concat(other)
		return n
	} else {
		other, _ := other.(*opNode)
		other.left = n.Concat(other.left)
		return other
	}
}

type node interface {
	Calc() float64
	Priority() int
	Concat(node) node
}
