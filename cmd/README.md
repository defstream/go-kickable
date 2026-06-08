# cani

`cani` is a command-line tool that answers the age-old question: can I kick it?

## Build

```shell
make build
```

## Usage

```shell
./bin/cani -kick "it"
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-kick` | `""` | The item to determine as kickable |

## Install system-wide

```shell
make install
```

This copies `bin/cani` to `$GOPATH/bin`.
