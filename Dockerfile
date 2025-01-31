FROM registry.uniphore.com/chainguard/go:latest-dev AS dev

ENV GOPRIVATE github.com/uniphore,gitlab.com/uniphore
ENV GOFLAGS=-mod=vendor
ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GOOS linux

WORKDIR /app

COPY cmd cmd
COPY pkg pkg
COPY internal internal

COPY go.mod .
COPY go.sum .
COPY vendor vendor

RUN go build -tags nomsgpack -o api ./cmd/api/main.go

ENTRYPOINT ["./api"]

# Production image
FROM registry.uniphore.com/chainguard/static:latest AS app
# Use the following image if you need dynamic glibc (e.g. for Kafka)
# FROM registry.uniphore.com/chainguard/glibc-dynamic:latest AS app

WORKDIR /app

COPY --from=dev /app/api .

ENTRYPOINT ["./api"]
