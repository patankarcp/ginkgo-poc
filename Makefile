# Makefile for managing Ginkgo PoC monorepo

# Variables
REPO_NAME = ginkgo-poc
REGISTRY = 
VERSION = latest
KUBE_NAMESPACE = default

# Directories
MICROSERVICES_DIR = microservices
PKG_DIR = pkg
FLUENTD_DIR = fluentd
HELM_CHART_DIR = helm/ginkgo-poc

# Build Docker images
.PHONY: build
build: build-postgres build-redis build-client build-fluentd

build-postgres:
	docker build -t $(REGISTRY)/postgres:$(VERSION) $(MICROSERVICES_DIR)/postgres

build-redis:
	docker build -t $(REGISTRY)/redis:$(VERSION) $(MICROSERVICES_DIR)/redis

build-client:
	docker build -t $(REGISTRY)/client:$(VERSION) $(MICROSERVICES_DIR)/client

build-fluentd:
	docker build -t $(REGISTRY)/fluentd:$(VERSION) $(FLUENTD_DIR)

# Push Docker images
.PHONY: push
push: push-postgres push-redis push-client push-fluentd

push-postgres:
	docker push $(REGISTRY)/postgres:$(VERSION)

push-redis:
	docker push $(REGISTRY)/redis:$(VERSION)

push-client:
	docker push $(REGISTRY)/client:$(VERSION)

push-fluentd:
	docker push $(REGISTRY)/fluentd:$(VERSION)

# Run Go tests
.PHONY: test
test:
	go test -v ./...

# Deploy with Helm
.PHONY: deploy
deploy:
	helm upgrade --install $(REPO_NAME) $(HELM_CHART_DIR) --namespace $(KUBE_NAMESPACE) --values $(HELM_CHART_DIR)/values.yaml

# Clean Docker images
.PHONY: clean
clean:
	docker rmi $(REGISTRY)/postgres:$(VERSION) || true
	docker rmi $(REGISTRY)/redis:$(VERSION) || true
	docker rmi $(REGISTRY)/client:$(VERSION) || true
	docker rmi $(REGISTRY)/fluentd:$(VERSION) || true

# Start local Kubernetes cluster using Kind (if needed)
.PHONY: kind
kind:
	kind create cluster --name $(REPO_NAME)

# Delete local Kubernetes cluster
.PHONY: kind-delete
kind-delete:
	kind delete cluster --name $(REPO_NAME)

# Lint Helm charts
.PHONY: lint
lint:
	helm lint $(HELM_CHART_DIR)

# Default target
.PHONY: all
all: build test

