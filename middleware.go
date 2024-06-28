package sake

type Middleware func()

type MiddlewareContainer struct {
	middlewares []Middleware
}

func NewMiddlewareContainer() *MiddlewareContainer {
	return &MiddlewareContainer{
		middlewares: []Middleware{},
	}
}

func (mc *MiddlewareContainer) Register(m Middleware) {
	mc.middlewares = append(mc.middlewares, m)
}
