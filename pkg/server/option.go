package server

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
)

type Option interface{ apply(r *Server) }

func WithServerLogger(l Logger) Option { return serverLoggerOption{l} }

func WithServerTracer(t opentracing.Tracer) Option { return serverTracerOption{t} }

func WithServerConfig(c Config) Option { return serverConfigOption{c} }

func WithServerPort(p int) Option { return serverPortOption{p} }

func WithServerReadTimeout(t int) Option { return serverReadTimeoutOption{t} }

func WithServerWriteTimeout(t int) Option { return serverWriteTimeoutOption{t} }

func WithShutdownDelaySeconds(d int) Option { return serverShutdownDelaySecondsOption{d} }

func WithHealthCheck(f func(http.HandlerFunc) http.HandlerFunc) Option {
	return serverHealthCheckOption{f}
}

func WithLivenessCheck(f func(http.HandlerFunc) http.HandlerFunc) Option {
	return serverLivenessCheckOption{f}
}

func WithReadinessCheck(f func(http.HandlerFunc) http.HandlerFunc) Option {
	return serverReadinessCheckOption{f}
}

func WithServerRouter(r Handler) Option {
	return serverRouterOption{r: r}
}

type serverLoggerOption struct{ logger Logger }

func (l serverLoggerOption) apply(s *Server) {
	if l.logger != nil {
		s.logger = l.logger
	}
}

type serverTracerOption struct{ tracer opentracing.Tracer }

func (t serverTracerOption) apply(s *Server) {
	if t.tracer != nil {
		s.tracer = t.tracer
	}
}

type serverConfigOption struct{ config Config }

func (m serverConfigOption) apply(s *Server) {
	s.config = m.config
}

type serverPortOption struct{ port int }

func (p serverPortOption) apply(s *Server) {
	s.config.Port = p.port
}

type serverReadTimeoutOption struct{ t int }

func (p serverReadTimeoutOption) apply(s *Server) {
	s.config.ReadTimeoutMs = p.t
}

type serverWriteTimeoutOption struct{ t int }

func (p serverWriteTimeoutOption) apply(s *Server) {
	s.config.WriteTimeoutMs = p.t
}

type serverShutdownDelaySecondsOption struct{ t int }

func (p serverShutdownDelaySecondsOption) apply(s *Server) {
	s.config.ShutdownDelaySeconds = p.t
}

type serverReadinessCheckOption struct {
	f func(http.HandlerFunc) http.HandlerFunc
}

func (r serverReadinessCheckOption) apply(s *Server) {
	s.readinessCheck = r.f
}

type serverLivenessCheckOption struct {
	f func(http.HandlerFunc) http.HandlerFunc
}

func (l serverLivenessCheckOption) apply(s *Server) {
	s.livenessCheck = l.f
}

type serverHealthCheckOption struct {
	f func(http.HandlerFunc) http.HandlerFunc
}

func (h serverHealthCheckOption) apply(s *Server) {
	s.healthCheck = h.f
}

type serverRouterOption struct {
	r Handler
}

func (sro serverRouterOption) apply(s *Server) {
	if sro.r != nil {
		s.Router = sro.r
	}
}

type FactoryOption interface{ apply(p *factory) }

func WithLogger(l Logger) FactoryOption { return loggerOption{logger: l} }

func WithTracer(t opentracing.Tracer) FactoryOption { return tracerOption{tracer: t} }

func WithConfig(c Config) FactoryOption { return configOption{c} }

func WithRouter(rf func() Handler) FactoryOption { return routerOption{rf} }

type tracerOption struct{ tracer opentracing.Tracer }

func (t tracerOption) apply(f *factory) {
	if t.tracer != nil {
		f.tracer = t.tracer
	}
}

type loggerOption struct{ logger Logger }

func (l loggerOption) apply(f *factory) {
	if l.logger != nil {
		f.logger = l.logger
	}
}

type configOption struct{ c Config }

func (co configOption) apply(f *factory) {
	f.config = co.c
}

type routerOption struct{ rf func() Handler }

func (ro routerOption) apply(f *factory) {
	if ro.rf != nil {
		f.routerFunc = ro.rf
	}
}
