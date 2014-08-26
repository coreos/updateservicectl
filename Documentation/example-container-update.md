# Example Update Workflow with Containers

Getting a distributed application running across a CoreOS cluster is [easy with fleet]({{site.url}}/docs/launching-containers/launching/fleet-example-deployment/). The hard part is deploying an update to the containers that make up the system, without having any downtime. The update service provides the tools we need to make this happen in a graceful and transparent way.

This orchestration happens through an updater which is responsible for reporting the progress of the update ([via the Omaha protocol](https://github.com/coreos/updateservicectl/blob/master/Documentation/protocol.md)) and executing the commands that actually verify and apply the update. Each application can have their own updater to execute app-specific commands or you can share the same update logic across the applications that you run. This guide will implement an extremely simple updater that uses the `updateservicectl watch` command to be notified of an update. It's important to note that using this scheme shouldn't be used for any production system.

To orchestrate the update, we need to run a "sidekick" for each of the app containers, which will check for an update and run a few [fleetctl]({{site.url}}/docs/launching-containers/launching/fleet-using-the-client/) commands if there is a new version.

We're going to build off the example described in the [Zero Downtime Frontend Deploys with Vulcand]() blog post. This guide is going to assume your application is able to function with mixed versions running. Before continuing, follow the directions in the blog post for scenario 1, a rolling frontend update. Check that vulcan is serving traffic correctly but stop when you reach the "Start Version 2.0.0" step. Instead of launching those manually, we're going to automatically update them.

## Create an Application, Channel and Package

Before we can update the example application, we need to set it up in the update service. You can read the detailed information in the [Getting Started]() guide, or quickly execute these commands:

```
$ updateservicectl app update --app-id=e96281a6-d1af-4bde-9a0a-97b76e56dc57 \
	--label="Example App" --description="Sample application to test container updates"
$ updateservicectl group create --app-id=e96281a6-d1af-4bde-9a0a-97b76e56dc57 \
	--channel=master --group-id=13DEFFFC-90C2-4B94-B3C2-1322BE8DC4E3 \
	--label="Example Group"
$ updateservicectl package create --app-id=e96281a6-d1af-4bde-9a0a-97b76e56dc57 \
	--version=1.0.0 --file=coreos/example:1.0.0
$ updateservicectl channel update --app-id=e96281a6-d1af-4bde-9a0a-97b76e56dc57 --channel=master --version=1.0.0
```

In the web interface, you should see everything that you just created:

[screenshot]

### Update Unit

For each instance of our example app already running, we need to run an updater as an additional sidekick. Our simple updater is a bash script that can be found in the [example update repository]().

Similar to the registration sidekick, the updater will use `X-ConditionMachineOf` will indicate to fleet that this container *must* run on the same machine as our primary container.

Since we're running multiple copies, each updater needs to reference a specific nginx instance. In order to use multiple instances of the same updater unit, we can use the variable `%i` to encode our reference into the filename. Naming our unit `updater@example-v1.0.0-A.service` will reference `example-v1.0.0-A.service`.

```
[Unit]
Description=Updater
After=docker.service
Requires=docker.service
# Binds this unit and the primary unit together. When the primary is stopped, this unit will be stopped too.
BindsTo=%i.service

[Service]
ExecStart=/usr/bin/docker run -name updater coreos/example-updater /bin/sh /update.sh
ExecStop=/usr/bin/docker stop updater

[X-Fleet]
X-ConditionMachineOf=%i.service
```

### Start The Units

Create three copies of the updater unit and start them:

```
fleetctl start updater@example*
```

## Verify Running Containers

There are three different ways we can verify that the units came up correctly. First, using fleet, we can `list-units`. You should see the three units running on the same machine as the primary unit:

```
$ fleetctl list-units
INSERT ME
```

Second, open the load balancer in your browser and you should see version 1.0.0 being printed out. Refresh a few times in order to be routed to different containers.

By now the updaters should have reported back to the update service. Browse the the application and group you set up earlier. You should see 3 instances running for that application.

[screenshot]

### Troubleshooting the Updater

If you run into problems with the updater, the best way to diagnose issues is to read the journal. Using the `-follow` flag will continually display new log lines as they are generated:

```
fleetctl journal -follow updater@example-v1.0.0-A.service.service
```

For more journalctl information, check out the [Reading the System Log]({{site.url}}/docs/cluster-management/debugging/reading-the-system-log/) guide.


## Trigger Update

Now that we know everything is running, let's load a new package into the update service, configure the update behavior for the group, and update the channel to the new package version.

### Load the 2.0.0 Package

Create a new package pointing to the 2.0.0 tag of the container:

```
updateservicectl package create --app-id=e96281a6-d1af-4bde-9a0a-97b76e56dc57 \
	--version=2.0.0 --file=coreos/example:2.0.0
```

### Start the Update

In order to inspect our cluster mid-update, you can slow down the roll-out rate. Set the group to update 1 instance every 5 minutes. Once we promote the 2.0.0 package, we should see only one container update.

Now we're ready to update. In the UI, modify the master channel to point to the 2.0.0 package that we loaded in. Since our updater is checking so frequently, we should see the new instance show up in the graph:

[screenshot]

We've successfully updated one of our containers! In your web browser, refresh the load balancer a few times. You should see a few requests coming from the updated container. Over the next 15 minutes, all of the containers should update.

## Further Reading

If you're interested in writing a more robust updater with customized logic for your application, check out the [protocol](protocol.md) documentation.
