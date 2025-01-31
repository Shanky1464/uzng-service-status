APP_NAME := $(shell yq .name charts/hello-world-go/Chart.yaml)

OSNAME := $(shell uname -s)
OSARCH := $(shell uname -m)

export GOPRIVATE=github.com/uniphore,gitlab.com/uniphore
export GOFLAGS=-mod=vendor
export CGO_ENABLED 0
export GOARCH amd64


.PHONY: vendor
vendor:
	go mod vendor

.PHONY: build
build:
ifeq ($(OS), Windows_NT)
	$(eval OS := windows)
else
    ifeq ($(OSNAME), Darwin)
		$(eval OS := mac)

    else ifeq ($(OSNAME), Linux)
		$(eval OS := linux)
    else
		$(eval OS := UNKNOWN)
    endif
endif

	@$(MAKE) build-$(OS)

.PHONY: build-linux
build-linux: vendor
	env GOOS=linux go build -o $(APP_NAME) ./cmd/api/main.go

.PHONY: build-mac
build-mac: vendor
ifeq ($(OSARCH), x86_64)
	$(eval OSARCH := amd64)
endif

	env GOOS=darwin GOARCH=$(OSARCH) go build -o $(APP_NAME) ./cmd/api/main.go

.PHONY: build-windows
build-windows: vendor
	env GOOS=windows go build -o $(APP_NAME) .\cmd\api\main.go

.PHONY: test-unit
test-unit:
	go test -v -short -tags=unit -covermode=count -coverprofile=coverage_unit.out ./...
	go tool cover -html coverage_unit.out -o coverage_unit.html

.PHONY: test-integration
test-integration:
	go test -v -short -p=1 -tags=integration -covermode=count -coverprofile=coverage_integration.out ./...
	go tool cover -html coverage_integration.out -o coverage_integration.html

.PHONY: test
test: test-unit test-integration

.PHONY: mocks
mocks:
	rm -rf pkg/mocks
	mockery --dir=pkg --output=pkg/mocks --all --keeptree

.PHONY: updatedeps
updatedeps:
	go get -v -u ./...
	go mod tidy

.PHONY: clean
clean:
	go clean -i ./...
	rm -f coverage_unit.* coverage_integration.*
