# Using the Client

`updatectl` lets you control and test the CoreOS update service. Subcommands
let you manage users, groups, packages and write a very simple client that gets
its state via environment variables.

## Update Clients

There are two tools to test out the update service: fakeclient and watch. Fake
client simulates a number of clients from a single command. Watch is used to
quickly implement a simple update client with a minimal amount of code.

### Fake Clients

This example will start 132 fake clients pinging the update service every 1 to
50 seconds against the CoreOS application's UUID and put them in the beta group
starting at version 1.0.0.

```
./bin/updatectl fakeclients -c 132 -m 1 -M 50 e96281a6-d1af-4bde-9a0a-97b76e56dc57 beta 1.0.0
```

### Update Watcher

The fakeclient is useful for generating traffic but if you want a fast way to
create your own client you can use `watch`. This will exec a program of your
choosing every time a new update is available.

First, create a simple application that dumps the environment variables that
the watcher will pass in. Call the script `updater.sh`.

```
#!/bin/sh
env | grep UPDATE_SERVICE
```

Next we will generate a random client id with and start watching for changes to the given app:

```
./bin/updatectl watch e96281a6-d1af-4bde-9a0a-97b76e56dc57 beta $(uuidgen) ./updater.sh
```

If you change the version of the beta group's channel then your script will be
re-executed and you will see the UPDATE_SERVICE environment variables change.

## Administative Flags

There are a few flags that you must provide to the administrative commands below. 

- `-u` is your username, usually this is an email address or `admin`
- `-k` is your API key
- `-s` is the URL to your update service instance

The commands below will all have a prefix like this:

```
./bin/updatectl -u admin -k d3b07384d113edec49eaa6238ad5ff00 -s https://example.update.core-os.net
```

## Application Management

Applications have two pieces of data: a universal unique identifier (UUID) and
a name. The UUID is used by all of the updaters to tell the service what
application they belong to.

### Add an Application

Create an application called CoreOS using its UUID along with a nice description.

```
./bin/updatectl update-app e96281a6-d1af-4bde-9a0a-97b76e56dc57 "CoreOS" "Linux for Servers"
```

### List Applications

```
./bin/updatectl list-apps
```

## Package Management

Packages represent an individual version of an application and the URL
associated with it. You can also include metadata like cryptographic hashes and
size for verification purposes.

### Add an Application Version

Pass in all of the metadata you could ever want to a new package.

```
./bin/updatectl new-package e96281a6-d1af-4bde-9a0a-97b76e56dc57 1.0.5 \
	--name foobar \
	--path update.gz \
	--size 23 \
	--sha1sum fe7374bddde2ddf07f6bfcc728d115d14338964b \
	--sha256sum b602d630f0a081840d0ca8fc4d35810e42806642b3127bb702d65c3df227d0f5 \
	--signature ixi6Oebo \
	--metadata-size 190
```

### List Application Versions

```
./bin/updatectl list-packages e96281a6-d1af-4bde-9a0a-97b76e56dc57
```

## Channel Management

A channel gives a nice symbolic name to packages. A group tracks a specific
channel. Think of channels as a DNS name for a package.

### Update a channel

A channel has a version of individual applications. To change the version of an
application specify the app id, channel and the version that channel
should present.

```
./bin/updatectl update-channel e96281a6-d1af-4bde-9a0a-97b76e56dc57 master 1.0.1
```

## Group Management

Clients get their updates by giving the service a combination of their group
and application id. Groups are usually some division of data centers,
environments or customers.

### Creating a Group

Create a group for the CoreOS application pointing at the master channel called
testing. This group might be used in your test environment.

```
./bin/updatectl new-group e96281a6-d1af-4bde-9a0a-97b76e56dc57 master testing "Testing Group"
```

### Pausing Updates on a Group

```
./bin/updatectl pause-group e96281a6-d1af-4bde-9a0a-97b76e56dc57 testing
```

### List Groups

```
./bin/updatectl list-groups
Label           Token                                   UpdatesPaused
Default Group   default                                 false
```

## Client Management

The service keeps track of clients and gives you a number of tools to see their
state. Most of these endpoints are more nicely consumed via the control panel
but you can use them from `updatectl` too.

### List Clients

This will list all clients that have been seen since the given timestamp.

```
./bin/updatectl list-updateclients --start 1392401442
```

This will list the clients grouped by AppId and Version

```
./bin/updatectl list-appversions --start 1392401442
```

### List Application Versions Over Time

```
./bin/updatectl list-eventversion --start 1392884048
Version Timestamp       Count
1.0.21  1392884400      20
1.0.24  1392884460      20
1.0.24  1392884520      20
```

## User management

### Create a new user

```bash
./bin/updatectl admin-create-user -u 'user@coreos.net'
```
