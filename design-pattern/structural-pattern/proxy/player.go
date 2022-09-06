package proxy

import "fmt"

type Player struct {
	champion Champion
	isBanned bool
}

func NewPlayer(champion Champion) *Player {
	return &Player{
		champion: champion,
	}
}

func (p *Player) Ban() {
	p.isBanned = true
}

func (p *Player) Unban() {
	p.isBanned = false
}

func (p Player) Q() {
	if !p.isBanned {
		p.champion.Q()
	} else {
		fmt.Println("Your account is banned! Cant use Q")
	}
}

func (p Player) W() {
	if !p.isBanned {
		p.champion.W()
	} else {
		fmt.Println("Your account is banned Cant use W")
	}
}

func (p Player) E() {
	if !p.isBanned {
		p.champion.E()
	} else {
		fmt.Println("Your account is banned Cant use E")
	}
}

func (p Player) R() {
	if !p.isBanned {
		p.champion.R()
	} else {
		fmt.Println("Your account is banned Cant use R")
	}
}
