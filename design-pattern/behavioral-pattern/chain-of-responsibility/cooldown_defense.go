package chain_of_responsibility

import (
	"time"
)

type CooldownDefense struct {
	PersistentDefense

	cooldown int
	lastUsed int64
}

func (cd CooldownDefense) IsReady() bool {
	now := time.Now().Unix()

	return now-cd.lastUsed > int64(cd.cooldown)
}
