package subsumption_test

import (
	. "."
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
