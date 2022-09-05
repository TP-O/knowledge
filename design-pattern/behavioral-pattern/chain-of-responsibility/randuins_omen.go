package chain_of_responsibility

import "fmt"

type RanduinsOmen struct {
	PersistentDefense
}

func NewRanduinsOmen() Defensible {
	return &RanduinsOmen{
		PersistentDefense: PersistentDefense{
			name: "Randuin's Omen",
		},
	}
}

func (ro RanduinsOmen) Defense(attack *Attack) {
	if attack.Type() == NormalAttack {
		var maxHealthOfOwner float32 = 2500
		var reduction float32 = 5.0 + 3.5*(maxHealthOfOwner/1000)

		attack.Reduce(int(reduction))

		fmt.Println(ro.name, "reduced", reduction, "damage")

		if attack.Damage() <= 0 {
			return
		}
	} else {
		fmt.Println(ro.name, "can't reduce damage of this attack")
	}

	ro.PersistentDefense.Defense(attack)
}
