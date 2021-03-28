# Cases

## Get Document

- `Client` sends a query to `WebSvr`
- if query contains a shortLink, query from `CacheSvr` at first
- if `CacheSvr` returns a doc, returns it
- otherwise, query `DB` for results
  - if query contains a shortLink, cache the result


## Post Document

- `Client` posts a document to `WebSvr`
- `WebSvr` saves the document to `DB`

