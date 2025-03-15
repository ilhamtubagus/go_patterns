package server

type Options struct {
	MaxConnections int
	TLS            bool
	ID             string
}

func defaultOptions() Options {
	return Options{
		MaxConnections: 100,
		TLS:            false,
		ID:             "default",
	}
}

type OptFunc func(*Options)

type Server struct {
	Options
}

func NewServer(opts ...OptFunc) *Server {
	op := defaultOptions()
	for _, fn := range opts {
		fn(&op)
	}

	return &Server{
		Options: op,
	}
}

func WithTLS(opts *Options) {
	opts.TLS = true
}

func WithMaxConnections(maxConn int) OptFunc {
	return func(options *Options) {
		options.MaxConnections = maxConn
	}
}

func WithID(id string) OptFunc {
	return func(options *Options) {
		options.ID = id
	}
}
