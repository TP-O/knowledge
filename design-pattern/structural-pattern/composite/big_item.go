package composite

import "fmt"

type BigItem struct {
	name              string
	blacksmithingCost int
	smallItems        []Item
}

func (bi BigItem) Price() int {
	price := bi.blacksmithingCost

	fmt.Println("**")

	for _, item := range bi.smallItems {
		price += item.Price()
	}

	fmt.Println("=>", bi.name, ":", price, "with a blacksmithing cost of", bi.blacksmithingCost)
	fmt.Println("**")

	return price
}

func (bi *BigItem) AddMaterial(item Item) {
	bi.smallItems = append(bi.smallItems, item)
}
