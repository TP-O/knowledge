package composite

type Kindlegem struct {
	BigItem
}

func NewKindlegem(cost int) *Kindlegem {
	return &Kindlegem{
		BigItem{
			name:              "Kindlegem",
			blacksmithingCost: cost,
		},
	}
}
