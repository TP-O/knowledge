package chain_of_responsibility

type Defensible interface {
	Defense(attack *Attack)
	SetNext(defense Defensible)
}
