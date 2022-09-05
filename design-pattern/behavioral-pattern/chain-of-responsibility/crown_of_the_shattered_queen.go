package chain_of_responsibility

import (
	"fmt"
	"time"
)

type CrownOfTheShatteredQueen struct {
	CooldownDefense
}

func NewCrownOfTheShatteredQueen() Defensible {
	return &CrownOfTheShatteredQueen{
		CooldownDefense{
			PersistentDefense: PersistentDefense{
				name: "Crown of the Shattered Queen",
			},
			cooldown: 40,
		},
	}
}

func (cotsq *CrownOfTheShatteredQueen) Defense(attack *Attack) {
	if cotsq.IsReady() {
		reduction := 0.75 * float32(attack.Damage())

		attack.Reduce(int(reduction))
		cotsq.lastUsed = time.Now().Unix()

		fmt.Println(cotsq.name, "reduced", reduction, "damage")
	} else {
		fmt.Println(cotsq.name, "can't reduce damage of this attack")
	}

	cotsq.PersistentDefense.Defense(attack)
}
