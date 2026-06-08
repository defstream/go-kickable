# Kite Service

Exposes the kickable API using [Kite](https://github.com/koding/kite), Koding's RPC-over-WebSocket framework. The server registers a `CanIKick` method accessible via the Kite protocol.

## Build

```shell
make kite
```

Produces `bin/cani.kite.server` and `bin/cani.kite.client`.

## Server

```shell
./bin/cani.kite.server -address localhost:3000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `localhost:3000` | Address to listen on |

## Client

```shell
./bin/cani.kite.client -kick "it" -addr http://127.0.0.1:3000/kite
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |
| `-addr` | `http://127.0.0.1:3000/kite` | Kite server URL |
