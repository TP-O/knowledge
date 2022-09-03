package facade

type WomboCombo struct {
	mf *MissFortune
	y  *Yasuo
	m  *Malphite
	k  *Kathus
	g  *Galio
}

func NewWomboCombo() *WomboCombo {
	return &WomboCombo{
		mf: &MissFortune{},
		y:  &Yasuo{},
		m:  &Malphite{},
		k:  &Kathus{},
		g:  &Galio{},
	}
}

func (wc *WomboCombo) PerfectWomboCombo() {
	wc.m.R()
	wc.y.R()
	wc.g.R()
	wc.mf.R()
	wc.k.R()
}

func (wc *WomboCombo) FailedWomboCombo() {
	wc.m.R()

	println("2 hours later...")

	wc.y.R()
}
