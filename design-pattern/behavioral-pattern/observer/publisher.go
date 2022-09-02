package observer

type Publisher[T any] interface {
	Subscribe(s Subscriber[T])

	Unsubscribe(s Subscriber[T])

	NotifySubscribers(notification T)
}
