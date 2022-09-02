package observer

type Subscriber[T any] interface {
	Update(data T)
}
