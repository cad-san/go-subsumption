package subsumption

type Agent struct {
	behaviors []Behavior
}

type Behavior interface {
	Init()
	Sence() (bool, error)
	Perform() (bool, error)
}

func (self Agent) Size() int {
	return len(self.behaviors)
}
