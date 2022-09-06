package proxy

type Yasuo struct {
	qStacks uint
}

func (y *Yasuo) Q() {
	if y.qStacks < 2 {
		println("Yasuo: Steel Tempest")

		y.qStacks++
	} else {
		println("Yasuo: Steel Tempest with a whirlwind")

		y.qStacks = 0
	}
}

func (y Yasuo) W() {
	println("Yasuo: Wind Wall")
}

func (y Yasuo) E() {
	println("Yasuo: Sweeping Blade")
}

func (y Yasuo) R() {
	println("Yasuo: Last Breath")
}
