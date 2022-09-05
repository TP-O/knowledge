package chain_of_responsibility

import "fmt"

func Run() {
	yp := NewYasuoPassive()
	crown := NewCrownOfTheShatteredQueen()
	heart := NewFrozenHeart()
	omen := NewRanduinsOmen()
	veil := NewBansheesVeil()
	jack := NewEdgeOfNight()

	defense := veil
	defense.SetNext(jack)
	jack.SetNext(crown)
	crown.SetNext(heart)
	heart.SetNext(omen)
	omen.SetNext(yp)

	skillAttack01 := &Attack{
		types:  SkillAttack,
		damage: 500,
	}
	skillAttack02 := &Attack{
		types:  SkillAttack,
		damage: 500,
	}
	skillAttack03 := &Attack{
		types:  SkillAttack,
		damage: 500,
	}

	// normalAttack01 := &Attack{
	// 	types:  NormalAttack,
	// 	damage: 10,
	// }
	// normalAttack02 := &Attack{
	// 	types:  NormalAttack,
	// 	damage: 10,
	// }
	// normalAttack03 := &Attack{
	// 	types:  NormalAttack,
	// 	damage: 500,
	// }

	fmt.Println("===Skill attack 01 with", skillAttack01.damage, "damage")
	defense.Defense(skillAttack01)
	fmt.Println("===Damage taken:", skillAttack01.Damage())
	fmt.Println()

	fmt.Println("===Skill attack 02 with", skillAttack02.damage, "damage")
	defense.Defense(skillAttack02)
	fmt.Println("===Damage taken:", skillAttack02.Damage())
	fmt.Println()

	fmt.Println("===Skill attack 03 with", skillAttack03.damage, "damage")
	defense.Defense(skillAttack03)
	fmt.Println("===Damage taken:", skillAttack03.Damage())
	fmt.Println()

	// fmt.Println("===Normal attack 01 with", normalAttack01.damage, "damage")
	// defense.Defense(normalAttack01)
	// fmt.Println("===Damage taken:", normalAttack01.Damage())
	// fmt.Println()

	// fmt.Println("===Normal attack 02 with", normalAttack02.damage, "damage")
	// defense.Defense(normalAttack02)
	// fmt.Println("===Damage taken:", normalAttack02.Damage())
	// fmt.Println()

	// fmt.Println("===Normal attack 03 with", normalAttack03.damage, "damage")
	// defense.Defense(normalAttack03)
	// fmt.Println("===Damage taken:", normalAttack03.Damage())
	// fmt.Println()
}
