package calc

// minus number must wrap with parentheses
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
