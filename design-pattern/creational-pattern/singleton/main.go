package singleton

import "fmt"

func Run() {
	db1 := Instance()
	db1.SetName("DB1")

	db2 := Instance()
	db2.SetName("DB2")

	fmt.Println("db1's name:", db1.Name())
	fmt.Println("db2's name:", db2.Name())
}
