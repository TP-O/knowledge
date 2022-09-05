package memento

import "fmt"

func Run() {
	i1 := NewItem(1, 500)
	i2 := NewItem(2, 1000)
	i3 := NewItem(3, 700)

	s := NewInventoryState(2000)
	i := NewInventory(s)

	fmt.Println("Init")
	i.Show()

	fmt.Println("Shopping...")
	fmt.Println("===")
	i.Add(i1)
	i.Add(i2)
	i.Add(i3)

	fmt.Println("===")
	fmt.Println("Check invetory")
	i.Show()

	fmt.Println("Undo")
	i.Undo()

	fmt.Println("===")
	fmt.Println("Check invetory")
	i.Show()

	fmt.Println("Buy item 3")
	fmt.Println("===")
	i.Add(i3)

	fmt.Println("===")
	fmt.Println("Check invetory")
	i.Show()
}
