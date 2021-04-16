package crawl

type config struct {
	target      string
	count       int
	upperID     int
	concurrency int
}

type configFn func(*config)

func WithTarget(target string) configFn {
	return func(c *config) {
		c.target = target
	}
}

func WithCount(count int) configFn {
	return func(c *config) {
		c.count = count
	}
}

func WithUpperID(upperID int) configFn {
	return func(c *config) {
		c.upperID = upperID
	}
}

func WithConcurrency(concurrency int) configFn {
	return func(c *config) {
		c.concurrency = concurrency
	}
}
