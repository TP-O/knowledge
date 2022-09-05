package chain_of_responsibility

import "fmt"

type FrozenHeart struct {
	PersistentDefense
}

func NewFrozenHeart() Defensible {
	return &FrozenHeart{
		PersistentDefense: PersistentDefense{
			name: "Frozen Heart",
		},
	}
}

func (fh FrozenHeart) Defense(attack *Attack) {
	if attack.Type() == NormalAttack {
		var maxHealthOfOwner float32 = 2500
		var reduction float32 = 7.0 + 3.5*(maxHealthOfOwner/1000)

		attack.Reduce(int(reduction))

		fmt.Println(fh.name, "reduced", reduction, "damage")

		if attack.Damage() <= 0 {
			return
		}
	} else {
		fmt.Println(fh.name, "can't reduce damage of this attack")
	}

	fh.PersistentDefense.Defense(attack)
}
