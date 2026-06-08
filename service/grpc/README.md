# gRPC Service

Exposes the kickable API over [gRPC](https://grpc.io/). Supports optional TLS via certificate and key flags.

## Build

```shell
make grpc
```

Produces `bin/cani.grpc.server` and `bin/cani.grpc.client`.

## Regenerate protobuf

Requires `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc`.

```shell
make proto-grpc
```

## Server

```shell
./bin/cani.grpc.server -address :4000
```

With TLS:

```shell
./bin/cani.grpc.server -address :4000 -crt server.crt -key server.key
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `:4000` | Address to listen on |
| `-crt` | `""` | Path to TLS certificate |
| `-key` | `""` | Path to TLS private key |

## Client

```shell
./bin/cani.grpc.client -kick "it" -address :4000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |
| `-address` | `:4000` | Server address |
