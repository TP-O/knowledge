package bridge

func Run() {
	y := &Yasuo{}
	c1 := NewController(y)

	c1.QQQR()

	println()
	println("===Leesin===")
	println()

	l := &Lessin{}
	c2 := NewController(l)

	c2.QQWR()
}
