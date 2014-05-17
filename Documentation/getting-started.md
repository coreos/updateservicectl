# Getting Started with the Update Service

The update service is a tool that helps you manage large-scale rolling upgrades of software. The service consists of three main parts:

1. A distributed web application that runs on top of fleet and docker.
2. `updatectl` a CLI interface to the service
3. Communication specification for your applications to report their current status and receive notifications of an available update.

The update service is based on [Omaha](https://code.google.com/p/omaha/), an open protocol from Google. This protocol powers updates for the Chrome browser, ChromeOS, Google Earth and more.

## Accessing the Update Service

The update service is an optional hosted service provided by CoreOS and is not included in a standard CoreOS cluster. Head over to the [Update Service]() page for more details.

Authentication for `updatectl` is done with a username and API key combination. Additional users and API keys can be provisioned by an existing user. Substitute the server address you were given during the activation process:

```
updatectl -u admin -k d3b07384d113edec49eaa6238ad5ff00 -s https://example.update.core-os.net <command>
```

Since you'll have to provide these flags each time, it's recommended that you set up an alias in your bash profile. We'll assume that you've done this for the rest of this document:

```
alias updatectl="/bin/updatectl -u admin -k d3b07384d113edec49eaa6238ad5ff00 -s https://example.update.core-os.net"
```

## Anatomy of an Update

Let's walk through the different parts of an update then use `updatectl` to simulate the release of an update.

### Application

You can use the update service to facilitate the roll-out of a new version of any application. An application is made up of group of instances. Each instance reports a unique identifier, version, group ID and status to the application via an updater.

You can view the current list of applications with `updatectl list-apps`:

```
$ updatectl list-apps
```

### Group

Applications contain groups of instances that are related. Groups are designed to be flexible and can reflect your company's preferred organizational scheme. Common eamples include groups by business unit, team, location, or environment. Each group has separate settings such as the ability to roll-out updates at a specific rate, be paused or unpaused and track a specific channel of the application.

For example, if multiple teams have deployed a distributed database, each team can control how
quickly a new version is rolled out based on their specific needs. Updates can also be paused per group
if a team doesn't want any updates.

```
updatectl list-groups <appid>
```

### Channel

Each application can specify channels, such as alpha or beta, that can be updated to refer to different packages. Channels allow you to upload a new beta package and have it rolled out to all groups that track the beta channel, with one command.

```
updatectl list-channels <appid>
```

### Updater

The updater is a small application that implements the update protocol. This code normally runs inside a separate container running "beside" the container of the application that is being updated. The updater is responsible for periodically reporting information about the instance. If an update is available, the response will contain a location to fetch the update package from plus information required to validate the package.

All of the logic that controls how the release is rolled-out lives in the update service in order for the updaters to be as simple as possible. The updater's job is fairly straightforward:

| Action | Description |
|--------|-------------|
| Send Request | Send current information about the instance to the update service. |
| Parse Response | Parse the response from the update service and determine any action that needs to take place. |
| Download and Verify Package | Download the package from the specified location and verify its signatures. |
| Send Progress | As a package is being downloaded or installed, send download progress and status to the update service. |
| Execute Update | Install the package and verify that it installed correctly. |
| Execute Healthcheck | Carry out a health check on the application to make sure it was installed and initialized correctly. |
| Send Update Complete | Once a package is successfully installed, tell the update service that everything was successful. |
| Send Update Error | If an error occurs during the package download or verification, or during update eexecution or health check, report an error. |

### Package

A package is a pointer to a signed blob of data that is stored somewhere accessible by your application instances. Each package is tied to an application ID and a semantic version number.

When an out-of-date app instance is notified of an update, it is told where to find the data and how to validate that data. Each application's updater can contain customized logic to download and apply the update.

For example, a webapp container could be stored in a private docker registry and the update response contains the full address to the registry and a specific tag to pull: `index.example.com/webapp:1.2.1`. The updater would be programmed to `docker pull` the image. If everything goes well, the updater sends the update complete event to the update service informing it that the new version is running.

### Relationship Diagram

![Relationship Diagram](img/Relationships.png)

## Test Update with Fake Clients

The easiest way to illustrate how these concepts work together is to trigger an example with some simulated clients. You can do this with either the UI or via `updatectl`, which is what we're going to use.

### Create an Application, Channel and Group

First set up a new application with a unique identifier, name and description:

```
updatectl update-app e96281a6-d1af-4bde-9a0a-97b76e56dc57 "FakeApp" "Fake app for testing"
```

You should now see it in the list of apps:

```
$updatectl list-apps
INSERT ME
```

Next, create a channel that our group of fake clients will track. Let's call it `master` and start it out on version `1.0.0`:

```
updatectl update-channel e96281a6-d1af-4bde-9a0a-97b76e56dc57 master 1.0.0
```

Next, create a group that we'll associate our fake clients with. Be sure to include the app id, the `master` channel, a name and a friendly label. 

```
$ updatectl new-group e96281a6-d1af-4bde-9a0a-97b76e56dc57 master fake1 "Fake Clients"
```

### Uploading and Signing a Package

Now that we have our application, group and channel set up, we can almost test an upgrade. The last step is to load in a new package. In this example, the new package will be fake, with an incremented version.

Start by preparing a fake `update.gz` and fake `update.meta`:

```
touch update-1.1.0.gz
echo '{"metadata_size": "0", "metadata_signature_rsa": "xxx"}' > update-1.1.0.meta
``

You can now use the `new-package` command to publish this fake package as version `1.1.0` (with a fake URL):

```
updatectl new-package e96281a6-d1af-4bde-9a0a-97b76e56dc57 \
    --version 1.1.0 \
    --file update-1.1.0.gz \
    --meta update-1.1.0.meta \
    --url https://fakepackage.local/update-1.1.0.gz
```

### Start Fake Clients

`updatectl` contains a tool to help you simulate many fake clients running your applcation. We're going to start 10 fake clients that are checking for updates every 30-60 seconds. This is much much faster than usual but it will allow us to see our changes take place quickly.

When we start the fake clients, we don't expect them to do anything since we're already on version `1.0.0`. In another terminal window, start the clients:

```
$ updatectl fakeclients -c 10 -m 30 -M 60 e96281a6-d1af-4bde-9a0a-97b76e56dc57 master 1.0.0
{fake-client-000}: noupdate
{fake-client-002}: noupdate
{fake-client-001}: noupdate
...
```

### Roll-out the Update

Now let's see how the fake clients react when we promote the new package `1.1.0` to the master channel. First, let's set the rate limit of the group to slow down the roll-out. This will make it easier to see what's going on. Since we only have 10 clients, 2 updates per 60 seconds should be slow enough:

```
$ updatectl update-group e96281a6-d1af-4bde-9a0a-97b76e56dc57 aca12080-0e1c-4baa-82eb-945b7416d6cd --updateCount 2 --updateInterval 60
Fake Clients	e96281a6-d1af-4bde-9a0a-97b76e56dc57		aca12080-0e1c-4baa-82eb-945b7416d6cd	false	2	60
```

Next, promote our `1.1.0` release on the `master` channel:

```
$ updatectl update-channel e96281a6-d1af-4bde-9a0a-97b76e56dc57 master 1.1.0
1.1.0
```

In the terminal window running the fake clients, you should see a few of them start to upgrade. The output looks like:

```
{fake-client-000}: updated from 1.0.0 to 1.1.0
```

In the UI, navigate to the app and group, then click on "View All Graphs". You should see the instances slowly start to converge on version `1.1.0`:

## Further Reading

Now that you've got the basics, checkout the [example update workflow]() for a practical application of the update service. If you're ready to start writing a custom update client for your application, the [Omaha protocol spec]() is a good place to start. The complete list of update service docs can be [found here]().
