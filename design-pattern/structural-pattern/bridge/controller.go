package bridge

type Controller struct {
	champion Champion
}

func NewController(champion Champion) *Controller {
	return &Controller{
		champion: champion,
	}
}

func (c Controller) QQQR() {
	c.champion.Q()
	c.champion.Q()
	c.champion.Q()
	c.champion.R()
}

func (c Controller) QQWR() {
	c.champion.Q()
	c.champion.Q()
	c.champion.W()
	c.champion.R()
}
