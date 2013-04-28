package behaviour

import (
	"testing"
)

func add(x, y int) int {
	return x + y
}

func ExampleDescribe() {
	var t *testing.T // this will be passed as an argument in your testrunner
	Describe(t, "Addition Feature", func(b Behaviour) {
		b.It("Should add two negative numbers", func() {
			if -5 != add(-2, -3) {
				t.Error("-2 + -3 should be -5")
			}
		})
		b.It("Should add two positive numbers", func() {
			if 5 != add(2, 3) {
				t.Error("2 + 3 should be 5")
			}
		})
		b.It("Should add zero", func() {
			if -5 != add(0, -5) {
				t.Error("-5 + 0 should be -5")
			}
		})
	})
}

func TestRunner(t *testing.T) {
	Describe(t, "Addition Feature", func(b Behaviour) {
		b.It("Should add two negative numbers", func() {
			if -5 != add(-2, -3) {
				t.Error("-2 + -3 should be -5")
			}
		})
		b.It("Should add two positive numbers", func() {
			if 5 != add(2, 3) {
				t.Error("2 + 3 should be 5")
			}
		})
		b.It("Should add zero", func() {
			if -5 != add(0, -5) {
				t.Error("-5 + 0 should be -5")
			}
		})
		b.It("Should fail", func() {
			if true {
				t.Error("Don't worry, this failure is expected, just to test failure output")
			}
		})
	})
}
