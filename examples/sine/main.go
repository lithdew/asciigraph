package main

import (
	"fmt"
	"github.com/lithdew/asciigraph"
	"math"
)

func main() {
	data := make([]float64, 105)
	for i := 0; i < 105; i++ {
		data[i] = 15 * math.Sin(float64(i)*((math.Pi*4)/120.0))
	}

	fmt.Println(asciigraph.Plot(data, asciigraph.Height(15)))
}
