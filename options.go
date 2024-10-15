package stack

type options struct {
	skip     int
	depth    int
	allstack bool
}

type Option func(*options)

func WithSkip(skip int) Option {
	return func(o *options) {
		o.skip = skip
	}
}

func WithDepth(depth int) Option {
	return func(o *options) {
		o.depth = depth
		o.allstack = false
	}
}

func WithAllStack() Option {
	return func(o *options) {
		o.allstack = true
	}
}
