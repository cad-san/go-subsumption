package subsumption_test

import (
	. "."
	"testing"
)

func TestCreate(t *testing.T) {
	agent := Agent{}

	if agent.Size() != 0 {
		t.Errorf("agent.Size() should be 0")
	}
}
