package observer

import "golang.org/x/exp/slices"

type Manga struct {
	name       string
	subcribers []Subscriber[string]
}

func NewManga(name string) Publisher[string] {
	return &Manga{
		name:       name,
		subcribers: make([]Subscriber[string], 0),
	}
}

func (m *Manga) Subscribe(s Subscriber[string]) {
	m.subcribers = append(m.subcribers, s)
}

func (m *Manga) Unsubscribe(s Subscriber[string]) {
	removedIndex := slices.IndexFunc(m.subcribers, func(e Subscriber[string]) bool {
		return e == s
	})

	if removedIndex != -1 {
		m.subcribers = slices.Delete(m.subcribers, removedIndex, removedIndex+1)
	}
}

func (m *Manga) NotifySubscribers(nofication string) {
	for _, subscriber := range m.subcribers {
		subscriber.Update(m.name + " notifies that " + nofication)
	}
}
