# GraphQL Service

Exposes the kickable API as a [GraphQL](https://graphql.org/) endpoint at `/graphql`.

## Build

```shell
make graphql
```

Produces `bin/cani.graphql.server`.

## Server

```shell
./bin/cani.graphql.server -address :8000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `:8000` | Address to listen on |

## Query

Send a GraphQL query using any HTTP client:

```shell
curl -X POST http://localhost:8000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ canIKick(it: \"it\") }"}'
```
