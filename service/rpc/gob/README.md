# RPC/GOB Service

Exposes the kickable API using Go's standard `net/rpc` package with [gob](https://pkg.go.dev/encoding/gob) encoding.

## Build

```shell
make rpc-gob
```

Produces `bin/cani.gob.rpc.server` and `bin/cani.gob.rpc.client`.

## Server

```shell
./bin/cani.gob.rpc.server -address :4000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `:4000` | Address to listen on |

## Client

```shell
./bin/cani.gob.rpc.client -kick "it" -address localhost:4000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |
| `-address` | `localhost:4000` | Server address |
