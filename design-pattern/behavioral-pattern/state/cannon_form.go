package state

import "fmt"

type CannonForm struct {
	champion *Jayce
}

func (cf CannonForm) Q() {
	fmt.Println("Shock Blast")
}

func (cf CannonForm) W() {
	fmt.Println("Hyper Charge")
}

func (cf CannonForm) E() {
	fmt.Println("Accerleration Gate")
}

func (cf CannonForm) R() {
	cf.champion.ChangeForm(HammerForm{champion: cf.champion})

	fmt.Println()
	fmt.Println("Switch to Mercury Hammer")
	fmt.Println("")
}
