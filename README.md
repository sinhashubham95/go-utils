# Go Utils - Fully Tested, Benchmarked and Error Free Set of Utilities

[![tag](https://img.shields.io/github/tag/sinhashubham95/go-utils.svg)](https://github.com/sinhashubham95/go-utils/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.13-%23007d9c)
[![GoDoc](https://godoc.org/github.com/sinhashubham95/go-utils?status.svg)](https://pkg.go.dev/github.com/sinhashubham95/go-utils)
![Build Status](https://github.com/sinhashubham95/go-utils/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/sinhashubham95/go-utils)](https://goreportcard.com/report/github.com/sinhashubham95/go-utils)
[![codecov](https://codecov.io/gh/sinhashubham95/go-utils/branch/master/graph/badge.svg?token=D6ND052IDH)](https://codecov.io/gh/sinhashubham95/go-utils)
[![Contributors](https://img.shields.io/github/contributors/sinhashubham95/go-utils)](https://github.com/sinhashubham95/go-utils/graphs/contributors)
[![License](https://img.shields.io/github/license/sinhashubham95/go-utils)](./LICENSE)

✨ **`sinhashubham95/go-utils` is an End-To-End Golang Utilities library compatible with Go version >= 1.18. You name it you find it here.**

This library provides a ready replacement for most of the Golang standard packages, and also offers many more valuable abstractions. It's completely light-weight implementing all the methods natively, and not bloating this beautiful library with any additional dependencies.

## 🚀 Install

```sh
go get github.com/sinhashubham95/go-utils@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

This library has no dependencies outside the Go standard library.

## 💡 Usage

You can import `go-utils` package using the following code snippet:

```go
package sample

import (
    "fmt"
    "github.com/sinhashubham95/go-utils/strings"
)
```

Then use one of the helpers below:

```go
func sample() {
	fmt.Prinln(strings.Join('n', 'a', 'r', 'u', 't', 'o'))
}
```

## 🤠 Spec

GoDoc: [https://godoc.org/github.com/sinhashubham95/go-utils](https://godoc.org/github.com/sinhashubham95/go-utils)
