# RPC/JSON Service

Exposes the kickable API using Go's standard `net/rpc` package with [jsonrpc](https://pkg.go.dev/net/rpc/jsonrpc) encoding (JSON-RPC 1.0).

## Build

```shell
make rpc-json
```

Produces `bin/cani.json.rpc.server` and `bin/cani.json.rpc.client`.

## Server

```shell
./bin/cani.json.rpc.server -address :4000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `:4000` | Address to listen on |

## Client

```shell
./bin/cani.json.rpc.client -kick "it" -address :4000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |
| `-address` | `:4000` | Server address |
