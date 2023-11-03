package go_best_practice

// Server 结构体
type Server struct {
	Host string
	Port int
}

// ServerOptions 结构体用来暂存所有的选项
type ServerOptions struct {
	Host string
	Port int
}

// Option 定义一个函数类型, 该函数接受一个指向 ServerOptions 的指针并修改它.
type Option func(*ServerOptions)

// WithHost 是一个 Option 函数, 用于设置 ServerOptions 的 Host 字段.
func WithHost(host string) Option {
	return func(o *ServerOptions) {
		o.Host = host
	}
}

// WithPort 是一个 Option 函数, 用于设置 ServerOptions 的 Port 字段.
func WithPort(port int) Option {
	return func(o *ServerOptions) {
		o.Port = port
	}
}

// NewServer 是创建 Server 实例的工厂函数, 可以接受任意数量的 Option 函数.
func NewServer(opts ...Option) *Server {
	// 使用默认值初始化选项
	options := &ServerOptions{
		Host: "127.0.0.1",
		Port: 80,
	}

	// 应用所有 Option 函数来修改选项
	for _, opt := range opts {
		opt(options)
	}

	// 基于最终的选项创建 Server 实例
	server := &Server{
		Host: options.Host,
		Port: options.Port,
	}

	return server
}
