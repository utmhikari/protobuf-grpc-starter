# protobuf-grpc-starter

a tiny go project to show the feature & application of protobuf & grpc

通过一个golang的应用来展示protobuf+grpc相关功能以及应用

# Pre-Requisities

- `golang` v1.16 (latest builtin libraries)
- `protoc` for protobuf3, to compile protobuf files
- `protoc-gen-go` for protoc, to generate `pb` files in `golang`
- `protoc-gen-go-grpc` for protoc, to generate `grpc` files in `golang`
- `make` for processing `Makefile` script

# Description

## Structure

`Protobuf` is widely used in modern backend services for internal communication via `GRPC`

This project, inspired by [Pastebin](https://pastebin.com/), implements minimal functions of the application

- user posts a text, `WebSvr` returns the shortLink
- user can use shortLink as query/param to search a specific text, which may already be cached in `CacheSvr`
- user can also use keyword/author to search specific texts, but these results won't be cached
- shortLink query results should be cached in `CacheSvr`, which implements an in-memory LRU cache, at each query
  - `WebSvr` calls `GetDocument` & `SetDocument` functions on `CacheSvr` via keep-alive `GRPC` connection
  - `GetDocument`
    - `Client` sends a query via http request to `WebSvr`
    - if query contains a shortLink, query from `CacheSvr` by `GetDocument` rpc call
      - if `CacheSvr` returns a doc, returns it
      - otherwise, query `DB` for the result
        - `DB` stores data in a single file
        - if exists in `DB`, cache the result to `CacheSvr` by `SetDocument` rpc call
    - if query does not contain a shortLink, query `DB` for results
  - `SetDocument`
    - `Client` posts a document to `WebSvr`
    - `WebSvr` saves the document to `DB`
  
## Run Project

- Run `make` to make protos and `CacheSvr` & `WebSvr` binaries
- start `./bin/cachesvr`
- start `./bin/websvr`
  - http port of `WebSvr` should be the `externalPort` of `websvr` in `cluster.json`
  - `WebSvr` uses [gin](https://github.com/gin-gonic/gin) as engine, see `internal/svr/websvr/main/websvr.go` for route definitions
  - you can use `Postman` to mock http requests to `WebSvr`

## References

- protobuf
  - [protobuf-github](https://github.com/golang/protobuf)
  - [protobuf-go-tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)
  - [protobuf-language-guide](https://developers.google.com/protocol-buffers/docs/proto3)
  - [protobuf-examples](https://github.com/protocolbuffers/protobuf/tree/master/examples)
- grpc
  - [grpc-go-github](https://github.com/grpc/grpc-go)
  - [grpc-go](https://grpc.io/docs/languages/go/)

