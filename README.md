# About

`zube-cli` is a small Command Line Utility written in [go](https://go.dev) for interacting with [Zube](https://zube.io)

## Setup

### Build

Make sure you have `go` and `make` installed, then run

```bash
make
```

`zube` expects a configuration file with your `client_id` (located at `~/config/zube/config.yml`)

You can find out how to get the `client_id` in the [offical Zube docs](https://zube.io/docs/api#authentication-section).

Example of a `config.yml`:

```yaml
client_id: some-super-long-client-id
```

This file is also used to cache your `access token`, so make sure only you have access to it.

## Usage

Simply call `zube` to see a list of available commands and flags.

For example, to view your current user's information and check that everything works:

```bash
$ zube currentPerson

Username: Daniils-Petrovs
Name: Daniils Petrovs
```
