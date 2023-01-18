# Services with Service Registry

## Goal

The goal of this lab is to show how to use API gateway and service registry with load balancing and tracing.

## Service Registry

[Consul](https://developer.hashicorp.com/consul) is used as a service registry.

## API Gateway

[Fabio](https://fabiolb.net) is used as API gateway.

## Tracing

[Jaeger](https://www.jaegertracing.io) is used as a tracing system.

## Services

This project includes:

1. Product service
2. Stock service

They use:

- [Fiber](https://github.com/gofiber/fiber) as a web framework
- [Zap](https://github.com/uber-go/zap) as a logger

Those services do not have clear structure, instead the main goal is to show how to use service registry with load balancing and tracing.

## Pre-requisites

### Installation

```bash
go mod download
```

Download software:

- [Consul](https://developer.hashicorp.com/consul/downloads)
- [Fabio](https://github.com/fabiolb/fabio/releases)
- [Jaeger](https://www.jaegertracing.io/download/):

### Run

In separate terminals:

```bash
consul agent -dev
```

```bash
fabio
```

```bash
jaeger-all-in-one
```

### Build

```bash
make build
```

### Dev requirements

- Development reload tool: [reflex](https://github.com/cespare/reflex)

```bash
go install github.com/cespare/reflex@latest
```

### Run in dev mode

```bash
make watch
```
