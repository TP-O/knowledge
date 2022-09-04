package composite

import "fmt"

type SmallItem struct {
	name  string
	price int
}

func (si SmallItem) Price() int {
	fmt.Println(si.name, ":", si.price)

	return si.price
}
