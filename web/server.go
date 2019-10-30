package web

import (
	"github.com/kataras/iris"
)

type Component func(app *iris.Application)
type ServerConfig struct {
	Listen string
}

type Server struct {
	config ServerConfig

	proxy *iris.Application
}

func New(config ServerConfig) (handler *Server) {

	s := Server{
		config: config,
		proxy:  iris.New(),
	}
	return &s
}

func (s *Server) RegisterComponent(component Component) *Server {
	component(s.proxy)
	return s
}

func (s *Server) AdvancedConfig(handler func(app *iris.Application)) *Server {

	if handler != nil {
		handler(s.proxy)
	}
	return s
}

func (s *Server) SetHomePage(indexHtml string) *Server {
	s.proxy.Get("/", func(c iris.Context) {
		_, _ = c.HTML(indexHtml)
	})
	return s
}

func (s *Server) OnErrorCode(code int, onErr func(ctx iris.Context)) *Server {
	s.proxy.OnErrorCode(code, onErr)
	return s
}

func (s *Server) Start(components ...Component) {

	if len(components) > 0 {
		for _, component := range components {
			s.RegisterComponent(component)
		}
	}
	_ = s.proxy.Run(iris.Addr(s.config.Listen))
}
