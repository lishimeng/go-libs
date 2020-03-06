package web

import (
	"context"
	"github.com/kataras/iris"
	"net/http"
)

type Component func(app *iris.Application)
type ServerConfig struct {
	Listen string
}

type Server struct {
	config        ServerConfig
	delegate      *multiHandlerServer
	primaryProxy  *iris.Application
	primaryPath   string
	primarySchema string
}

const (
	SchemaHttp    = "http"
	SchemaHttps   = "https"
	DefaultSchema = SchemaHttp
)

func New(config ServerConfig) (handler *Server) {

	s := Server{
		config:        config,
		primaryProxy:  iris.New(),
		delegate:      newServer(),
		primaryPath:   "/",
		primarySchema: DefaultSchema,
	}
	return &s
}

func (s *Server) RegisterComponent(component Component) *Server {
	component(s.primaryProxy)
	return s
}

func (s *Server) AdvancedConfig(handler func(app *iris.Application)) *Server {

	if handler != nil {
		handler(s.primaryProxy)
	}
	return s
}

func (s *Server) SetHomePage(indexHtml string) *Server {
	s.primaryProxy.Get("/", func(c iris.Context) {
		_, _ = c.HTML(indexHtml)
	})
	return s
}

func (s *Server) OnErrorCode(code int, onErr func(ctx iris.Context)) *Server {
	s.primaryProxy.OnErrorCode(code, onErr)
	return s
}

func (s *Server) RegisterComponents(components ...Component) *Server {

	if len(components) > 0 {
		for _, component := range components {
			s.RegisterComponent(component)
		}
	}
	return s
}

func (s *Server) AddHttpHandler(schema string, pathGroup string, handler http.Handler) *Server {
	if s.delegate != nil {
		s.delegate.RegisterHandler(schema, pathGroup, handler)
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.primaryProxy.Build(); err != nil {
		return err
	}
	s.AddHttpHandler(s.primarySchema, s.primaryPath, s.primaryProxy)
	srv := http.Server{
		Addr:    s.config.Listen,
		Handler: s.delegate,
	}
	_ = srv.Shutdown(ctx)
	return srv.ListenAndServe()
}
