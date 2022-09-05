package chain_of_responsibility

type AttackType = int

const (
	NormalAttack AttackType = iota
	SkillAttack
)

type Attack struct {
	types  AttackType
	damage int
}

func NewAttack(types AttackType, damage int) *Attack {
	return &Attack{
		types:  types,
		damage: damage,
	}
}

func (a Attack) Type() AttackType {
	return a.types
}

func (a Attack) Damage() int {
	return a.damage
}

func (a *Attack) Reduce(damage int) {
	a.damage -= damage
}
