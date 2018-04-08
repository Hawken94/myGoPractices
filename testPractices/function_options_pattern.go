package testPractices

// 初始化结构体,默认参数

// TODO: 尽快看一下!!!
var defaultStuffClientOptions = StuffClientOptions{
	Retries: 3,
	Timeout: 2,
}

type StuffClientOption func(*StuffClientOptions)
type StuffClientOptions struct {
	Retries int
	Timeout int
}

func WithRetries(r int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Retries = r
	}
}
func WithTimeout(t int) StuffClientOption {
	return func(o *StuffClientOptions) {
		o.Timeout = t
	}
}

type Connection struct{}
type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}

func NewStuffClient(conn Connection, opts ...StuffClientOptions) stuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return &stuffClient{
		conn:    conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}
