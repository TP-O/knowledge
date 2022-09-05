package chain_of_responsibility

import (
	"fmt"
	"time"
)

type BansheesVeil struct {
	CooldownDefense
}

func NewBansheesVeil() Defensible {
	return &BansheesVeil{
		CooldownDefense{
			PersistentDefense: PersistentDefense{
				name: "Banshee's Veil",
			},
			cooldown: 40,
		},
	}
}

func (bv *BansheesVeil) Defense(attack *Attack) {
	if bv.IsReady() && attack.Type() == SkillAttack {
		attack.Reduce(attack.Damage())
		bv.lastUsed = time.Now().Unix()

		fmt.Println(bv.name, "blocked this attack")

		return
	}

	fmt.Println(bv.name, "can't block this attack")

	bv.PersistentDefense.Defense(attack)
}
