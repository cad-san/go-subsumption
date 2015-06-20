package subsumption

type Agent struct {
	behaviors []Behavior
}

type Behavior interface {
	Init()
	Sense() (bool, error)
	Perform() (bool, error)
}

func (self *Agent) Init() {
	for _, b := range self.behaviors {
		b.Init()
	}
}

func (self *Agent) Perform() (bool, error) {
	var active *Behavior = nil

	for _, b := range self.behaviors {
		if ret, err := b.Sense(); err != nil {
			continue
		} else if ret == true {
			active = &b
		}
	}

	if active == nil {
		return false, nil
	}

	return (*active).Perform()
}

func (self *Agent) AddBehavior(new_behavior Behavior) {
	self.behaviors = append(self.behaviors, new_behavior)
}

func (self Agent) Size() int {
	return len(self.behaviors)
}
