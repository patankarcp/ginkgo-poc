package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

type Factory interface {
	Create(options ...Option) *Server
}

type factory struct {
	tracer     opentracing.Tracer
	logger     Logger
	config     Config
	routerFunc func() Handler
}

func NewFactory(options ...FactoryOption) Factory {
	f := &factory{
		tracer:     opentracing.NoopTracer{},
		logger:     NoopLogger{},
		config:     defaultConfig(),
		routerFunc: func() Handler { return &http.ServeMux{} },
	}

	for _, option := range options {
		if option != nil {
			option.apply(f)
		}
	}

	return f
}

type Server struct {
	Router         Handler
	tracer         opentracing.Tracer
	logger         Logger
	config         Config
	livenessCheck  func(http.HandlerFunc) http.HandlerFunc
	readinessCheck func(http.HandlerFunc) http.HandlerFunc
	healthCheck    func(http.HandlerFunc) http.HandlerFunc
}

func (f *factory) Create(options ...Option) *Server {
	srvr := &Server{
		tracer: f.tracer,
		logger: f.logger,
		config: f.config,
		Router: f.routerFunc(),
	}

	for _, option := range options {
		if option != nil {
			option.apply(srvr)
		}
	}

	srvr.Router.HandleFunc("/live", srvr.getLivenessHandler())
	srvr.Router.HandleFunc("/ready", srvr.getReadinessHandler())
	srvr.Router.HandleFunc("/health", srvr.getHealthCheckHandler())

	return srvr
}

func (s *Server) Serve(ctx context.Context) error {
	handler := s.getHandler(ctx)
	port := s.config.Port
	if port < 1 {
		port = 8080
	}

	srvr := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  time.Duration(s.config.ReadTimeoutMs) * time.Millisecond,
		WriteTimeout: time.Duration(s.config.WriteTimeoutMs) * time.Millisecond,
	}

	errs := make(chan error)
	go func() {
		if err := srvr.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Error(ctx, "server failed to start up", "error", err)
			errs <- err
		} else {
			errs <- nil
		}
	}()

	s.logger.Info(ctx, "server started successfully", "port", port)

	go func() {
		errs <- s.gracefulShutdown(ctx, &srvr)
	}()

	return <-errs
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.getHandler(context.Background()).ServeHTTP(w, r)
}

func (s *Server) profilingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			s.logger.Debug(r.Context(), "http path response time",
				"path", r.URL.EscapedPath(),
				"method", r.Method,
				"time", time.Since(start),
			)
		}
		return http.HandlerFunc(fn)
	}
}

func (s *Server) tracingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return nethttp.Middleware(s.tracer, next)
	}
}

func (s *Server) timeoutMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, time.Duration(s.config.RequestTimeoutSec)*time.Second, "timeout")
	}
}

func (s *Server) getHandler(ctx context.Context) http.Handler {
	var h http.Handler = s.Router
	h = s.timeoutMiddleware()(h)
	h = s.tracingMiddleware()(h)
	h = s.profilingMiddleware()(h)
	return h
}

func (s *Server) gracefulShutdown(ctx context.Context, server *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit
	s.logger.Info(ctx, "signal received", "signal", sig)

	timeout := time.Duration(s.config.ShutdownDelaySeconds) * time.Second

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {

		s.logger.Error(
			ctx,
			"Error while gracefully shutting down server, forcing shutdown because of error",
			"err", err)
		return err
	}
	s.logger.Info(ctx, "server exited successfully")
	return nil
}

func (s *Server) getLivenessHandler() http.HandlerFunc {
	dflt := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	if s.livenessCheck != nil {
		return s.livenessCheck(dflt)
	}
	return dflt
}

func (s *Server) getReadinessHandler() http.HandlerFunc {
	defult := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	if s.readinessCheck != nil {
		return s.readinessCheck(defult)
	}
	return defult
}

func (s *Server) getHealthCheckHandler() http.HandlerFunc {

	defult := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK!"))
	})

	if s.healthCheck != nil {
		return s.healthCheck(defult)
	}
	return defult
}

type Handler interface {
	http.Handler
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type Config struct {
	Port                 int
	ReadTimeoutMs        int
	WriteTimeoutMs       int
	RequestTimeoutSec    int
	ShutdownDelaySeconds int
}

func defaultConfig() Config {
	return Config{
		Port:                 8080,
		ReadTimeoutMs:        10000,
		WriteTimeoutMs:       10000,
		RequestTimeoutSec:    10,
		ShutdownDelaySeconds: 5,
	}
}
