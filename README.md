# About

[![Go](https://github.com/platogo/zube-cli/actions/workflows/go.yml/badge.svg)](https://github.com/platogo/zube-cli/actions/workflows/go.yml)

`zube-cli` is a small Command Line Utility written in [Go](https://go.dev) for interacting with the [Zube](https://zube.io) issue tracking system.

## Features

- Fetch various Zube resources
- Construct complex Zube queries
- Create cards same way you would using the web client
- Client-side request caching

## Setup

### Manual install

Make sure you have `go` and `make` installed, then run

```bash
make
make install
```

Optionally install ZSH completions with:

```bash
make completions.zsh
```

### With Nix

Clone this repository and run

```bash
nix-build
nix profile install ./result
```

### Configuration

`zube` expects a configuration file with your **client_id**.
`zube` looks for this configuration file, in order, in:

- `~/config/zube/config.yml`
- `$XDG_CONFIG_HOME/zube/config.yml`

You can find out how to get the `client_id` in the [offical Zube docs](https://zube.io/docs/api#authentication-section).

Example of a `config.yml`:

```yaml
client_id: some-super-long-client-id
```

Easiest way to set your `client_id` is by:

Creating the configuration file, with either

``` bash
touch ~/config/zube/config.yml
```

``` bash
touch $XDG_CONFIG_HOME/zube/config.yml
```

Initializing the config

```bash
zube config init
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

- Cards
  - [x] interactive creation
  - [ ] editing
  - [ ] movement
  - [ ] archiving
  - [ ] commenting
  - [ ] Card queries / search by text
- [ ] Homebrew formula
- [ ] **Zube Query Launguage (ZQL)** parser as alternative for command line flag filters
- [ ] Filter support by name instead of just by IDs
- [ ] Optionally dump response data as JSON
- [ ] `zubed` daemon for periodic update polling
- [x] Move `zube` functionality into dedicated `zube-go` library
- Internal
  - [x] request caching
  - [ ] smart automatic auth using browser cookie store access