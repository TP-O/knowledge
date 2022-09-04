package composite

import "fmt"

type BigItem struct {
	name              string
	blacksmithingCost int
	materials         []Item
}

func (bi BigItem) Price() int {
	price := bi.blacksmithingCost

	fmt.Println("**")

	for _, item := range bi.materials {
		price += item.Price()
	}

	fmt.Println("=>", bi.name, ":", price, "with a blacksmithing cost of", bi.blacksmithingCost)
	fmt.Println("**")

	return price
}

func (bi *BigItem) AddMaterial(item Item) {
	bi.materials = append(bi.materials, item)
}
