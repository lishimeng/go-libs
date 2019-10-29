package web

import (
	"github.com/kataras/iris"
)

type Component func(app *iris.Application)
type ServerConfig struct {
	Listen string
}
type Server interface {
	Start(components ...Component)
	RegisterComponent(component Component)
	SetHomePage(indexHtml string)
	OnErrorCode(code int, onErr func(ctx iris.Context))
	ConfigExtra(handler func(app *iris.Application))
}

type server struct {
	config ServerConfig

	proxy *iris.Application
}

func Init(config ServerConfig) (handler *Server) {

	s := server{
		config: config,
		proxy:  iris.New(),
	}
	var h Server = &s
	handler = &h
	return handler
}

func (s *server) RegisterComponent(component Component) {
	component(s.proxy)
}

func (s *server) ConfigExtra(handler func(app *iris.Application)) {

	if handler != nil {
		handler(s.proxy)
	}
}

func (s *server) SetHomePage(indexHtml string) {
	s.proxy.Get("/", func(c iris.Context) {
		_, _ = c.HTML(indexHtml)
	})
}

func (s *server) OnErrorCode(code int, onErr func(ctx iris.Context)) {
	s.proxy.OnErrorCode(code, onErr)
}

func (s *server) Start(components ...Component) {

	if len(components) > 0 {
		for _, component := range components {
			s.RegisterComponent(component)
		}
	}
	_ = s.proxy.Run(iris.Addr(s.config.Listen))
}
