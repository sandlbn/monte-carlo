package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func simulate(samples int) float64 {
	var inside float64

	for i := 0; i < samples; i++ {
		if x, y := rand.Float64(), rand.Float64(); (x*x + y*y) <= 1 {
			inside++

		}

	}

	return (inside / float64(samples) * 4)

}

func main() {
	runs := []struct {
		samples int
	}{
		{200},
		{300},
	}

	for {
		for _, run := range runs {
			fmt.Printf("pi=%f runs=%d\n", simulate(run.samples), run.samples)
			time.Sleep(3 * time.Millisecond)

		}

	}
}
