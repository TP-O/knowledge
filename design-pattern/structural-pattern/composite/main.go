package composite

func Run() {
	s := NewSheen(700)
	rc := NewRubyCrystal(400)
	ls1 := NewLongSword(350)
	ls2 := NewLongSword(350)

	cw := NewCaulfieldsWarhammer(400)
	cw.AddMaterial(ls1)
	cw.AddMaterial(ls2)

	k := NewKindlegem(400)
	k.AddMaterial(rc)

	ds := NewDevineSunderer(700)
	ds.AddMaterial(cw)
	ds.AddMaterial(s)
	ds.AddMaterial(k)

	println("\nPrice of Devine Sunderer:", ds.Price())
}
