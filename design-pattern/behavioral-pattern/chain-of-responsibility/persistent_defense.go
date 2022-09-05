package chain_of_responsibility

type PersistentDefense struct {
	name string
	next Defensible
}

func (pd *PersistentDefense) SetNext(next Defensible) {
	pd.next = next
}

func (pd PersistentDefense) Defense(attack *Attack) {
	if pd.next != nil {
		pd.next.Defense(attack)
	}
}
