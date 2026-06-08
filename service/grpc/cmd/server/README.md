# cani.grpc.server

gRPC server for the kickable API.

## Usage

```shell
./bin/cani.grpc.server -address :4000
```

With TLS:

```shell
./bin/cani.grpc.server -address :4000 -crt server.crt -key server.key
```

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `:4000` | Address to listen on |
| `-crt` | `""` | Path to TLS certificate |
| `-key` | `""` | Path to TLS private key |
