package subsumption_test

import (
	. "."
	"errors"
	gomock "github.com/golang/mock/gomock"
	"testing"
)

func TestCreate(t *testing.T) {
	agent := Agent{}

	if agent.Size() != 0 {
		t.Errorf("agent.Size() should be 0")
	}
}

func TestAddBehavior(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	behabior := NewMockBehavior(ctrl)
	agent := Agent{}

	agent.AddBehavior(behabior)

	if agent.Size() != 1 {
		t.Errorf("agent.Size() should be 1")
	}
}

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	behabior := NewMockBehavior(ctrl)
	behabior.EXPECT().Init()

	agent := Agent{}
	agent.AddBehavior(behabior)

	agent.Init()
}

func TestPerform(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	behabior := NewMockBehavior(ctrl)
	behabior.EXPECT().Sense().Return(true, nil)
	behabior.EXPECT().Perform().Return(true, nil)

	agent := Agent{}
	agent.AddBehavior(behabior)

	ret, err := agent.Perform()

	if err != nil {
		t.Errorf("unexpected agent.Perform() fail with : %v", err)
	}
	if ret != true {
		t.Errorf("agent.Perform() should be true because behabior is active")
	}
}

func TestSenceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	behabior := NewMockBehavior(ctrl)
	behabior.EXPECT().Sense().Return(false, errors.New("unknown error"))

	agent := Agent{}
	agent.AddBehavior(behabior)

	ret, err := agent.Perform()

	if err != nil {
		t.Errorf("unexpected agent.Perform() fail with : %v", err)
	}
	if ret != false {
		t.Errorf("agent.Perform() should be false because no behavior is active")
	}
}
