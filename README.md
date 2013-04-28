# Pitch

Sometimes you want to create a test suite with examples defined dynamically. For example, if you are using an API that returns recommended music, and you want to test whether each recommended music track has defined metadata, it may be nice to dynamically create test examples. This package provides a few helpers to group test examples and contexts.

# Usage

```go
package yourpackage

import (
    "testing"
    "ioutil"
    . "github.com/pranavraja/behaviour"
)

func TestRunner(t *testing.T) {
    Describe(t, "My awesome API", func(b Behaviour) {
        b.Context("Music recommendations metadata", func(b Behaviour) {
            recommendations, _ := myAwesomeRecommendations()
            for track := range recommendations {
                b.It(track.Name, func() {
                    if track.Artist == "" {
                        t.Error("track has no artist")
                    }
                })
            }
        })
    })
}
```

If the test fails, you will get output to `t.Log` like this:

    --- FAIL: TestRunner (0.00 seconds)
        behaviour.go:35: My awesome API
        behaviour.go:35:   Music recommendations metadata
        behaviour.go:28:     Feeling Good
        behaviour.go:28:     Ain't No Sunshine
        yourpackage_test.go:15: track has no artist
        behaviour.go:28:     Never Let Me Go
    FAIL
    exit status 1
    FAIL    _/home/user/go/yourpackage  0.004s

...indicating that "Ain't No Sunshine" was the failed example.

# Testing

`go test`. Note that there is an expected failure, just to test the logging output.


