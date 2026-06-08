# go-micro Service

Exposes the kickable API using [go-micro.dev/v4](https://github.com/go-micro/go-micro). The service uses mDNS for automatic service discovery — no external registry required.

## Build

```shell
make micro
```

Produces `bin/cani.micro.server` and `bin/cani.micro.client`.

## Server

```shell
./bin/cani.micro.server
```

The server registers itself as `Kickable` on the local mDNS network. No address flags are needed.

## Client

```shell
./bin/cani.micro.client -kick "it"
```

The client discovers the `Kickable` service automatically via mDNS.

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |

## Notes

Both server and client must be on the same local network for mDNS discovery to work. For cross-network deployments, go-micro supports pluggable registries (etcd, consul, etc.).
