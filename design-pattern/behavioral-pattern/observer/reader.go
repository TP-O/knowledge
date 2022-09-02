package observer

import "fmt"

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(data string) {
	fmt.Println(r.name + " has received: " + data)
}
