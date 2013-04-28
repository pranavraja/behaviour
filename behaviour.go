package behaviour

import (
	"strings"
	"testing"
)

type Behaviour struct {
	desc  string
	depth int
	t     *testing.T
}

// Returns the behaviour description, indented according to its depth.
func (behaviour Behaviour) String() string {
	return strings.Repeat("  ", behaviour.depth) + behaviour.desc
}

// Creates a nested context to group behaviours.
func (behaviour Behaviour) Context(desc string, block func(Behaviour)) {
	b := Behaviour{desc, behaviour.depth + 1, behaviour.t}
	behaviour.t.Log(b)
	block(b)
}

// Specifies a behaviour. In `block` you may use standard testing functions
func (behaviour Behaviour) It(desc string, block func()) {
	behaviour.t.Log(strings.Repeat("  ", behaviour.depth+1) + desc)
	block()
}

// Specifies a feature.
func Describe(t *testing.T, desc string, block func(Behaviour)) Behaviour {
	b := Behaviour{desc, 0, t}
	t.Log(b)
	block(b)
	return b
}
