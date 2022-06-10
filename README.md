# About

[![Go](https://github.com/platogo/zube-cli/actions/workflows/go.yml/badge.svg)](https://github.com/platogo/zube-cli/actions/workflows/go.yml)

`zube-cli` is a small Command Line Utility written in [go](https://go.dev) for interacting with the [Zube](https://zube.io) issue tracking system.

## Features

- Fetch various Zube resources
- Construct complex Zube queries
- Client-side request caching

## Setup

### Build & Install

Make sure you have `go` and `make` installed, then run

```bash
make
make install
```

`zube` expects a configuration file with your **client_id** (located at `~/config/zube/config.yml`)

You can find out how to get the `client_id` in the [offical Zube docs](https://zube.io/docs/api#authentication-section).

Example of a `config.yml`:

```yaml
client_id: some-super-long-client-id
```

You must also save your **Zube private key** `.pem` file in `~/.ssh/zube_api_key.pem`

This file is also used to cache your `access token`, so make sure only you have access to it.

## Usage

Simply call `zube` to see a list of available commands and flags.

For example, to view your current user's information and check that everything works:

```bash
$ zube currentPerson

Username: Daniils-Petrovs
Name: Daniils Petrovs
```

To list cards by a filter (e.g. by `status`), you would do:

```bash
$ zube card ls --status done

Number Title                                          Status
------ ---------------------------------------------- ----------
13260  Use matrix in github build actions...          done
13252  Fix export timestamp...                        done
```

## Contributing

Read [CONTRIBUTING](CONTRIBUTING.md)

## Roadmap

- [ ] Card interactive creation
- [ ] Card editing & state transition
- [ ] Card queries / search by text
- [ ] Homebrew formula
- [ ] **Zube Query Launguage (ZQL)** parser as alternative for command line flag filters
- [ ] Filter support by name instead of just by IDs
- [ ] Optionally dump response data as JSON
- [ ] `zubed` daemon for periodic update polling
- [ ] Move `zube` functionality into dedicated `zube-go` library
