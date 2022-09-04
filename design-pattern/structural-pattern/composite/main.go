package composite

func Run() {
	s := NewSheen(700)
	rc1 := NewRubyCrystal(400)
	rc2 := NewRubyCrystal(400)
	ls := NewLongSword(350)

	p := NewPhase(350)
	p.AddMaterial(rc1)
	p.AddMaterial(ls)

	k := NewKindlegem(400)
	k.AddMaterial(rc2)

	ds := NewDevineSunderer(700)
	ds.AddMaterial(p)
	ds.AddMaterial(s)
	ds.AddMaterial(k)

	println("\nPrice of Devine Sunderer:", ds.Price())
}
