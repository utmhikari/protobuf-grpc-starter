# protobuf-grpc-starter

a tiny project to show the feature of protobuf & grpc

通过一个小project来展示protobuf+grpc相关功能及特性

# Cases

## Get Document

- `Client` sends a query to `WebSvr`
- if query contains a shortLink, query from `CacheSvr` by `GetDocument` rpc call
  - if `CacheSvr` returns a doc, returns it
  - otherwise, query db for the result
    - if exists, cache the result to `CacheSvr` by `SetDocument` rpc call
- if query does not contain a shortLink, query `DB` for results

## Post Document

- `Client` posts a document to `WebSvr`
- `WebSvr` saves the document to `DB`


## References

- protobuf
  - [protobuf-github](https://github.com/golang/protobuf)
  - [protobuf-go-tutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)
  - [protobuf-language-guide](https://developers.google.com/protocol-buffers/docs/proto3)
  - [protobuf-examples](https://github.com/protocolbuffers/protobuf/tree/master/examples)
- grpc
  - [grpc-go-github](https://github.com/grpc/grpc-go)
  - [grpc-go](https://grpc.io/docs/languages/go/)

