# cf inspect

A `cf` CLI plugin for inspecting various metadata from your Cloud Foundry deployment.

Usage:

```
$ cf inspect <COMMAND> [args...]
```

## Install

This plugin is an early experiment and must be built from source.

- Ensure you have installed Go 1.7+
- Install the `dep` tool: `go get -u github.com/golang/dep/cmd/dep`
- Clone this repo:
   - `mkdir -p $(go env GOPATH)/src/github.com/zmb3`
   - `git clone https://github.com/zmb3/cfinspect $(go env GOPATH)/src/github.com/zmb3/cfinspect`
   - `cd $(go env GOPATH)/src/github.com/zmb3/cfinspect`
- Install dependencies: `dep ensure`
- Build: `go build`

At this point, you should have a `cfinspect` that can be registered with the CLI:

```
$ cf install-plugin ./cfinspect
$ cf inspect -h
NAME:
   inspect - Inspect various CF metadata

USAGE:
   inspect
	cf inspect <command> [args...]
```

## Commands

### droplet

`cf inspect droplet <APP_NAME>` will download the droplet for the specified app
and print a summary of its contents.

```
$ cf inspect droplet staticapp
Type   Name                                     [Size]
---------------------------------------------------
d     ./
f       ./staging_info.yml                       74
d     ./logs/
d     ./tmp/
d     ./app/
f       ./app/sources.yml                        133
d     ./app/nginx/
d     ./app/nginx/logs/
d     ./app/nginx/sbin/
f       ./app/nginx/sbin/nginx                   5864002
d     ./app/nginx/conf/
f       ./app/nginx/conf/nginx.conf              1413
f       ./app/nginx/conf/mime.types              2081
f       ./app/boot.sh                            125
f       ./app/start_logging.sh                   86
f       ./app/Staticfile                         0
f       ./app/.profile                           265
d     ./app/.profile.d/
f       ./app/.profile.d/staticfile.sh           1214
d     ./app/public/
f       ./app/public/README.md                   1975
f       ./app/public/index.html                  393
f       ./app/public/LICENSE                     11357
f       ./app/public/index.js                    650
f       ./app/public/styles.css                  63
f       ./app/public/config.json                 2100

-- Droplet saved to droplet-staticapp-20170612-085112.tgz
```
