package options

// Apply calls the given opt functions with generic options struct.
//
//	func New(o ...options.Option[MyOptions]) {
//		opts := options.Apply(nil, o...)
//	 	...
//	}
func Apply[O any](opts *O, options ...Option[O]) *O {
	if opts == nil {
		opts = new(O)
	}

	for _, o := range options {
		if o != nil {
			o(opts)
		}
	}

	return opts
}

// Options is a generic function that represents some option, e.g.
//
//	func WithSize(size int) options.Option[MyOptions] {
//		return func(opts *MyOptions) {
//			opts.size = size
//		}
//	}
type Option[O any] func(*O)
