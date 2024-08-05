package main

import (
	"net/http"

	ghttp "github.com/patankarcp/ginkgo-poc/pkg/http"
	"github.com/patankarcp/ginkgo-poc/pkg/kafka"
	"github.com/patankarcp/ginkgo-poc/pkg/server"
)

type UserService struct {
	ServerFactory server.Factory

	HTTPConfig         ghttp.Config
	HTTPClientProvider ghttp.Provider

	KafkaConfig kafka.Config
	KafkaClient kafka.Client
}

func (uServ *UserService) List(w http.ResponseWriter, r *http.Request) {
	httpClient := uServ.HTTPClientProvider.GetWrappedClient(uServ.HTTPConfig)
	_, _ = httpClient.Get("http://hello.com/")

	kw, _ := uServ.KafkaClient.Writer(r.Context(), uServ.KafkaConfig)
	_, _ = kw.Write(r.Context(), "apple", []byte("message"))
	w.WriteHeader(http.StatusNoContent)
}

func (uServ *UserService) Get(w http.ResponseWriter, r *http.Request) {
	httpClient := uServ.HTTPClientProvider.GetWrappedClient(uServ.HTTPConfig)
	_, _ = httpClient.Get("http://hello.com/")

	kw, _ := uServ.KafkaClient.Writer(r.Context(), uServ.KafkaConfig)
	_, _ = kw.Write(r.Context(), "apple", []byte("message"))
	w.WriteHeader(http.StatusNoContent)
}

func (uServ *UserService) Create(w http.ResponseWriter, r *http.Request) {
	httpClient := uServ.HTTPClientProvider.GetWrappedClient(uServ.HTTPConfig)
	_, _ = httpClient.Get("http://hello.com/")

	kw, _ := uServ.KafkaClient.Writer(r.Context(), uServ.KafkaConfig)
	_, _ = kw.Write(r.Context(), "apple", []byte("message"))
	w.WriteHeader(http.StatusNoContent)
}

func (uServ *UserService) Update(w http.ResponseWriter, r *http.Request) {
	httpClient := uServ.HTTPClientProvider.GetWrappedClient(uServ.HTTPConfig)
	_, _ = httpClient.Get("http://hello.com/")

	kw, _ := uServ.KafkaClient.Writer(r.Context(), uServ.KafkaConfig)
	_, _ = kw.Write(r.Context(), "apple", []byte("message"))
	w.WriteHeader(http.StatusNoContent)
}

func (uServ *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	httpClient := uServ.HTTPClientProvider.GetWrappedClient(uServ.HTTPConfig)
	_, _ = httpClient.Get("http://hello.com/")

	kw, _ := uServ.KafkaClient.Writer(r.Context(), uServ.KafkaConfig)
	_, _ = kw.Write(r.Context(), "apple", []byte("message"))
	w.WriteHeader(http.StatusNoContent)
}
