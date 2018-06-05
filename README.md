# updateservicectl - CoreOS update service CLI

`updateservicectl` lets you control and test the CoreOS update service. Subcommands
let you manage users, groups, packages and write a very simple client that gets
its state via environment variables.

## Compatibility with CoreUpdate versions

`updateservicectl v2.0.0` is compatible with `CoreUpdate v2.2.0` or newer.

`updateservicectl v1.4.0` is compatible with `CoreUpdate v2.1.1` or older.

## About the Update Service

The update service is a tool that helps you manage large-scale rolling upgrades of software. The service consists of three main parts:

1. A distributed web application that runs within docker containers.
2. `updateservicectl` a CLI interface to the service
3. Communication specification for your applications to report their current status and receive notifications of an available update.

## Getting Started

Once you have gained access to your update service installation, check out the [Getting Started guide](http://coreos.com/products/coreupdate/docs/latest/getting-started.html) that will walk you through configuration of your applications, groups and update settings.

## Building the Client

Major releases for all platforms are listed under the [Releases tab](https://github.com/coreos/updateservicectl/releases) on this repository. If you'd like to build your own client:

1. `./build` or `make` (depending on the version of `updateservicectl` you are building)
2. The client is now built. Use it with `./bin/updateservicectl <command>`

## Creating Releases

You can build a release of a specfic version by running
`scripts/build-release <git-tag-of-version>`.

If you are a CoreOS developer, you may bump the version with
`scripts/bump-version <version>`. You can do this and build a release
at the same time with `scripts/new-release`.

Example: `scripts/new-release 0.2.0+git`.

## Documentation

[Using the Client](http://coreos.com/products/coreupdate/docs/latest/updatectl-client.html) - Read about all of the supported commands in `updateservicectl`

[Protocol](http://coreos.com/products/coreupdate/docs/latest/update-protocol.html) - A technical document about the Omaha protocol
