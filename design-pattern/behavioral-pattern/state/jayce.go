package state

type Jayce struct {
	form Form
}

func NewJayce() *Jayce {
	return &Jayce{}
}

func (j Jayce) Q() {
	j.form.Q()
}

func (j Jayce) W() {
	j.form.W()
}

func (j Jayce) E() {
	j.form.E()
}

func (j Jayce) R() {
	j.form.R()
}

func (j *Jayce) ChangeForm(form Form) {
	j.form = form
}
