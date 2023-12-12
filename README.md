# g2-sdk-go-mock

## :warning: WARNING: g2-sdk-go-mock is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing `g2-sdk-go-mock` packages provide mock
[Go](https://go.dev/)
objects representing the Software Development Kit that wraps the
Senzing C SDK APIs.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/g2-sdk-go-mock.svg)](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/g2-sdk-go-mock)](https://goreportcard.com/report/github.com/senzing/g2-sdk-go-mock)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/g2-sdk-go-mock/blob/main/LICENSE)

[![gosec.yaml](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/gosec.yaml/badge.svg)](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/Senzing/g2-sdk-go-mock/actions/workflows/go-test-windows.yaml)

## Overview

The Senzing `g2-sdk-go-mock` packages enable Go programs to simulate calling Senzing library functions.
The `g2-sdk-go-mock` implementation of the
[g2-sdk-go](https://github.com/Senzing/g2-sdk-go)
interface does not require the Senzing C libraries to be installed.

Other implementations of the
[g2-sdk-go](https://github.com/Senzing/g2-sdk-go)
interface include:

- [g2-sdk-go-base](https://github.com/Senzing/g2-sdk-go-base) - for
  calling Senzing Go SDK APIs natively
- [g2-sdk-go-grpc](https://github.com/Senzing/g2-sdk-go-grpc) - for
  calling Senzing SDK APIs over [gRPC](https://grpc.io/)
- [go-sdk-abstract-factory](https://github.com/Senzing/go-sdk-abstract-factory) - An
  [abstract factory pattern](https://en.wikipedia.org/wiki/Abstract_factory_pattern)
  for switching among implementations

## Use

(TODO:)

## References

1. [Development](docs/development.md)
1. [Errors](docs/errors.md)
1. [Examples](docs/examples.md)
1. [Package reference](https://pkg.go.dev/github.com/senzing/g2-sdk-go-mock)
