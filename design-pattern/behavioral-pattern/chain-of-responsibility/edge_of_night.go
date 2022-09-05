package chain_of_responsibility

import (
	"fmt"
	"time"
)

type EdgeOfNight struct {
	CooldownDefense
}

func NewEdgeOfNight() Defensible {
	return &EdgeOfNight{
		CooldownDefense{
			PersistentDefense: PersistentDefense{
				name: "Edge Of Night",
			},
			cooldown: 40,
		},
	}
}

func (eon *EdgeOfNight) Defense(attack *Attack) {
	if eon.IsReady() && attack.Type() == SkillAttack {
		attack.Reduce(attack.Damage())
		eon.lastUsed = time.Now().Unix()

		fmt.Println(eon.name, "blocked this attack")

		return
	}

	fmt.Println(eon.name, "can't block this attack")

	eon.PersistentDefense.Defense(attack)
}
