package state

func Run() {
	jayce := NewJayce()
	intialForm := CannonForm{champion: jayce}
	jayce.ChangeForm(intialForm)

	jayce.Q()
	jayce.E()
	jayce.W()
	jayce.R()

	jayce.Q()
	jayce.W()
	jayce.E()
	jayce.R()

	jayce.Q()
}
