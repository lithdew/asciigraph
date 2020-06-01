package asciigraph

type Config struct {
	Width  int
	Height int
}

func load(opts ...Option) Config {
	var cfg Config
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	return cfg
}

type Option interface{ apply(c *Config) }

type (
	applyWidth  int
	applyHeight int
)

func (w applyWidth) apply(c *Config)  { c.Width = int(w) }
func (h applyHeight) apply(c *Config) { c.Height = int(h) }

func Width(width int) Option {
	if width < 0 {
		width = 0
	}
	return applyWidth(width)
}

func Height(height int) Option {
	if height < 0 {
		height = 0
	}
	return applyHeight(height)
}
