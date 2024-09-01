package config

type Config struct {
	MaxDepth  int
	UserAgent string
	Timeout   int
}

func New(depth int) *Config {
	return &Config{
		MaxDepth:  depth,
		UserAgent: "FireCrawl/1.0",
		Timeout:   30,
	}
}