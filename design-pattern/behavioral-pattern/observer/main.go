package observer

func Run() {
	m1 := NewManga("Manga 01")
	m2 := NewManga("Manga 02")

	r1 := NewReader("Reader 01")
	r2 := NewReader("Reader 02")
	r3 := NewReader("Reader 03")
	r4 := NewReader("Reader 04")

	m1.Subscribe(r1)
	m1.Subscribe(r2)
	m1.Subscribe(r3)
	m1.Subscribe(r4)

	m2.Subscribe(r3)
	m2.Subscribe(r4)

	m1.NotifySubscribers("new chapter is released!")

	println("=====")

	m2.NotifySubscribers("new chapter is released!")

	println("=====")

	m1.Unsubscribe(r3)
	m1.Unsubscribe(r4)

	m1.NotifySubscribers("manga is dropped!")
}
