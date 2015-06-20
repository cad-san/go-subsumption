package subsumption

type Agent struct {
	behaviors []Behavior
}

type Behavior interface {
	Init()
	Sence() (bool, error)
	Perform() (bool, error)
}

func (self *Agent) Init() {
	for _, b := range self.behaviors {
		b.Init()
	}
}

func (self *Agent) AddBehavior(new_behavior Behavior) {
	self.behaviors = append(self.behaviors, new_behavior)
}

func (self Agent) Size() int {
	return len(self.behaviors)
}
