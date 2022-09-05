package chain_of_responsibility

import (
	"fmt"
	"time"
)

type yasuoPassive struct {
	CooldownDefense
}

func NewYasuoPassive() Defensible {
	return &yasuoPassive{
		CooldownDefense{
			PersistentDefense: PersistentDefense{
				name: "Yasuo Passive",
			},
			cooldown: 20,
		},
	}
}

func (yp *yasuoPassive) Defense(attack *Attack) {
	if yp.IsReady() {
		armor := 200

		attack.Reduce(armor)
		yp.lastUsed = time.Now().Unix()

		fmt.Println(yp.name, "reduced", armor, "damage")

		if attack.Damage() <= 0 {
			return
		}
	} else {
		fmt.Println(yp.name, "can't reduce damage of this attack")
	}

	yp.PersistentDefense.Defense(attack)
}
