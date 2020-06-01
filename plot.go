package asciigraph

import (
	"github.com/mattn/go-runewidth"
	"math"
	"strconv"
)

func Plot(series []float64, opts ...Option) string {
	if len(series) == 0 {
		return ""
	}

	config := load(opts...)

	// preprocess series data
	// figure out min and max of series data
	// figure out height if height is not set
	// figure out float precision for axis labels

	min := float64(0)
	max := float64(0)

	for _, point := range series {
		if point < min {
			min = point
		}
		if point > max {
			max = point
		}
	}

	interval := math.Abs(max - min)

	if config.Height == 0 {
		if int(interval) <= 0 {
			config.Height = int(interval * math.Pow10(int(math.Ceil(-math.Log10(interval)))))
		} else {
			config.Height = int(interval)
		}
	}

	ratio := 1.0
	if interval != 0.0 {
		ratio = float64(config.Height-1) / interval
	}

	min2 := math.Round(min * ratio)
	max2 := math.Round(max * ratio)

	logMax := math.Log10(math.Max(math.Abs(max), math.Abs(min)))
	if min == 0.0 && max == 0.0 {
		logMax = -1.0
	}

	precision := 2

	if logMax < 0 {
		if math.Mod(logMax, 1) != 0 {
			precision = precision + int(math.Abs(logMax))
		} else {
			precision = precision + int(math.Abs(logMax)-1.0)
		}
	} else if logMax > 2 {
		precision = 0
	}

	maxNumLen := runewidth.StringWidth(strconv.FormatFloat(max, 'f', precision, 64))
	minNumLen := runewidth.StringWidth(strconv.FormatFloat(min, 'f', precision, 64))

	numLen := minNumLen
	if numLen < maxNumLen {
		numLen = maxNumLen
	}

	// linearly interpolate data to fit graph
	// interpolate to width - max number length - 2
	// the extra 2 runes is for a space and separator between the axis and lines

	cols := len(series) + numLen + 2
	if config.Width > 0 {
		series = interpolate(series, config.Width-numLen-2)
		cols = config.Width
	}

	// abs(max - min)

	rows := int(max2 - min2)
	if rows < 0 {
		rows = -rows
	}

	// include an extra column for line breaks
	plot := make([]rune, (cols+1)*(rows+1))

	// leave the last character as eof (rune=0)
	// have every end of the column be a line break
	// have every other character be a space

	for i := 0; i < len(plot)-1; i++ {
		if (i+1)%(cols+1) == 0 {
			plot[i] = '\n'
		} else {
			plot[i] = ' '
		}
	}

	// draw axis and labels

	for y := int(min2); y <= int(max2); y++ {
		val := float64(y)
		if rows > 0 {
			val = max - (float64(y-int(min2)) * interval / float64(rows))
		}

		label := []rune(strconv.FormatFloat(val, 'f', precision, 64))

		x := (y - int(min2)) * (cols + 1)
		o := numLen - len(label)

		esc := false
		for i := 0; i < len(label); i++ {
			if x+o+i >= len(plot) || plot[x+o+i] == '\n' || plot[x+o+i] == 0 {
				esc = true
				break
			}
			plot[x+o+i] = label[i]
		}
		if esc || x+numLen+1 >= len(plot) || plot[x+numLen+1] == '\n' {
			continue
		}

		if y == 0 {
			plot[x+numLen+1] = '┼'
		} else {
			plot[x+numLen+1] = '┤'
		}
	}

	if len(series) == 0 {
		return string(plot)
	}

	// plot lines

	i := (rows-int(math.Round(series[0]*ratio)-min2))*(cols+1) + numLen + 1
	if i >= 0 && i < len(plot) {
		plot[i] = '┼'
	}

	for x := 0; x < len(series)-1; x++ {
		y0 := int(math.Round(series[x]*ratio) - float64(int(min2)))
		y1 := int(math.Round(series[x+1]*ratio) - float64(int(min2)))

		x1 := (rows-y0)*(cols+1) + x + numLen + 2
		x2 := (rows-y1)*(cols+1) + x + numLen + 2

		if x1 < 0 || x2 < 0 || x1 >= len(plot) || x2 >= len(plot) {
			continue
		}

		if y0 == y1 {
			plot[x1] = '─'
		} else {
			if y0 > y1 {
				plot[x2] = '╰'
				plot[x1] = '╮'
			} else {
				plot[x2] = '╭'
				plot[x1] = '╯'
			}

			start := int(math.Min(float64(y0), float64(y1))) + 1
			end := int(math.Max(float64(y0), float64(y1)))

			for y := start; y < end; y++ {
				i := (rows-y)*(cols+1) + x + numLen + 2
				if i < 0 || i >= len(plot) {
					continue
				}
				plot[i] = '│'
			}
		}
	}

	return string(plot)
}

func linear(before, after, factor float64) float64 {
	return before + (after-before)*factor
}

func interpolate(data []float64, size int) []float64 {
	if size < 0 {
		return nil
	}
	res := make([]float64, 0, size)
	res = append(res, data[0])

	smoothFactor := float64(len(data)-1) / float64(size-1)

	for i := 1; i < size-1; i++ {
		smooth := float64(i) * smoothFactor
		before := math.Floor(smooth)
		after := math.Ceil(smooth)

		res = append(res, linear(data[int(before)], data[int(after)], smooth-before))
	}

	res = append(res, data[len(data)-1])

	return res
}
