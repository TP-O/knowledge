package state

import "fmt"

type HammerForm struct {
	champion *Jayce
}

func (hf HammerForm) Q() {
	fmt.Println("To The Skies!")
}

func (hf HammerForm) W() {
	fmt.Println("Lightning Field")
}

func (hf HammerForm) E() {
	fmt.Println("Thundering Blow")
}

func (hf HammerForm) R() {
	hf.champion.ChangeForm(CannonForm{champion: hf.champion})

	fmt.Println()
	fmt.Println("Switch to Mercury Cannon")
	fmt.Println()
}
