package calc

// minus number must wrap with parentheses
func Calc(expression string) (float64, error) {
	bs := []byte(expression)
	words := splitWord(bs)
	err := checkGrammer(words)
	if err != nil {
		return 0, err
	}
	tree := makeTree(words)
	if tree == nil {
		terr := CalcErr("Empty expression")
		return 0, &terr
	}
	return tree.Calc()
}
