package asciigraph_test

import (
	"github.com/lithdew/asciigraph"
	"github.com/mattn/go-runewidth"
	"github.com/stretchr/testify/require"
	"math"
	"strings"
	"testing"
	"testing/quick"
)

func TestNoCrashes(t *testing.T) {
	f := func(data []float64, w uint, h uint) bool {
		w = w % 512
		h = h % 512

		for i := 0; i < len(data); i++ {
			data[i] = math.Mod(data[i], 10)
		}

		if len(data) == 0 || w == 0 || h == 0 {
			return true
		}

		graph := asciigraph.Plot(data, asciigraph.Width(int(w)), asciigraph.Height(int(h)))

		found := false
		for _, line := range strings.Split(graph, "\n") {
			if runewidth.StringWidth(line) == int(w) {
				found = true
				break
			}
		}

		return found
	}

	require.NoError(t, quick.Check(f, &quick.Config{MaxCount: 1000}))
}

func BenchmarkPlot(b *testing.B) {
	data := make([]float64, 105)
	for i := 0; i < 105; i++ {
		data[i] = 15 * math.Sin(float64(i)*((math.Pi*4)/120.0))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		asciigraph.Plot(data, asciigraph.Width(80), asciigraph.Height(24))
	}
}
