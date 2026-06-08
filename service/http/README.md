# HTTP Service

Exposes the kickable API over plain HTTP. The server responds to `GET /<item>` requests and returns a plain-text answer.

## Build

```shell
make http
```

Produces `bin/cani.http.server` and `bin/cani.http.client`.

## Server

```shell
./bin/cani.http.server -address localhost:3000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-address` | `localhost:3000` | Address to listen on |

## Client

```shell
./bin/cani.http.client -kick "it" -address http://localhost:3000
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |
| `-address` | `http://localhost:3000` | Server address |
