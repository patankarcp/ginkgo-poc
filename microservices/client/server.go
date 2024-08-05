package main

import (
	"github.com/patankarcp/ginkgo-poc/pkg/server"
)

func NewServer(service UserService) *server.Server {
	s := service.ServerFactory.Create()
	s.Router.HandleFunc("/list", service.List)
	s.Router.HandleFunc("/get", service.Get)
	s.Router.HandleFunc("/create", service.Create)
	s.Router.HandleFunc("/update", service.Update)
	s.Router.HandleFunc("/delete", service.Delete)
	return s
}
