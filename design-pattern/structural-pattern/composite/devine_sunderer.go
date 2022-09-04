package composite

type DevineSunderer struct {
	BigItem
}

func NewDevineSunderer(cost int) *DevineSunderer {
	return &DevineSunderer{
		BigItem: BigItem{
			name:              "Devine Sunderer",
			blacksmithingCost: cost,
		},
	}
}
