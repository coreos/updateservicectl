# updatectl - CoreOS update service CLI

`updatectl` lets you control and test the CoreOS update service. Subcommands
let you manage users, groups, packages and write a very simple client that gets
its state via environment variables.

## About the Update Service

The update service is a tool that helps you manage large-scale rolling upgrades of software. The service consists of three main parts:

1. A distributed web application that runs within docker containers.
2. `updatectl` a CLI interface to the service
3. Communication specification for your applications to report their current status and receive notifications of an available update.

## Getting Started

Once you have gained access to your update service installation, check out the [Getting Started guide](Documentation/getting-started.md) that will walk you through configuration of your applications, groups and update settings.

## Building the Client

Major releases for all platforms are listed under the [Releases tab](https://github.com/coreos-inc/updatectl/releases) on this repository. If you'd like to build your own client:

1. `./build`
2. The client is now built. Use it with `./updatectl <command>`

## Documentation

[Using the Client](Documentation/client.md) - Read about all of the supported commands in `updatectl`

[Example Upgrade Workflow](Documentation/example-container-update.md) - Walk through updating a set of containers running on CoreOS

[Protocol](Documentation/protocol.md) - A technical document about the Omaha protocol
