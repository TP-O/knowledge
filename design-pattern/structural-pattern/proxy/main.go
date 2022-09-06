package proxy

func Run() {
	p := NewPlayer(&Yasuo{})

	p.Q()
	p.Q()
	p.Q()

	p.Ban()

	p.Q()
	p.R()
}
