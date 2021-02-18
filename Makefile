MODULE = $(shell go list -m)
BIN  = $(CURDIR)/bin
M = $(shell printf "\033[34;1m▶\033[0m")
GO      = go
IMG ?= beelzebubabhi/reverse-proxy:latest

.PHONY: all
all: deps lint build

# Run go fmt against code
fmt:
	go fmt ./...

.PHONY: build
build: deps | $(BIN) ; $(info $(M) building executable…) @ ## Build program binary
	$(GO) build \
		-mod=vendor \
		-o $(BIN)/$(shell basename $(MODULE)) main.go


.PHONY: run
run:
	echo "Run using the binary. Generate the binary by running go build"

deps:
	#go clean --modcache
	go mod tidy
	go mod vendor

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) installing $(PACKAGE)…)
	   GOBIN=$(BIN) $(GO) install -mod=vendor $(PACKAGE)

GOLANGCI-LINT = $(BIN)/golangci-lint
$(BIN)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/cmd/golangci-lint

GOIMPORTS =$(BIN)/goimports
$(BIN)/goimports: PACKAGE=golang.org/x/tools/cmd/goimports


format: | $(GOIMPORTS); $(info $(M) Runing goimports formatter on all files except vendor)
		 @$(GOIMPORTS) -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

lint: format | $(GOLANGCI-LINT); $(info $(M) running golangci-lint) @ ## Run golangci-lint
			$(GOLANGCI-LINT) run ./...

# Build the docker image
docker-build: 
	docker build . -t ${IMG} --build-arg GIT_USER_NAME=${GIT_USER_NAME} --build-arg GIT_AUTH_KEY=${GIT_AUTH_KEY}

# Push the docker image
docker-push:
	docker push ${IMG}