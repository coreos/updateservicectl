# Example Update Workflow with Containers

Getting a distributed application running across a CoreOS cluster is [easy with fleet]({{site.url}}/docs/launching-containers/launching/fleet-example-deployment/). The hard part is deploying an update to the containers that make up the system, without having any downtime. The update service provides the tools we need to make this happen in a graceful and transparent way.

To orchestrate the update, we need to run a "sidekick" for each of the app containers, which will check for an update and run a few [fleetctl]({{site.url}}/docs/launching-containers/launching/fleet-using-the-client/) commands if there is a new version. It's important to note that the updater used in this example is extremely simple and shouldn't be used for any production system. 

To get started, let's take a closer look at the units that make up our application. If you don't have a CoreOS cluster running with at least 3 nodes, now would be a good time to [set one up]({{site.url}}/docs). All of the Dockerfiles and scripts used in this example can be found on [Github]().

## Create an Application, Channel and Package

Before we deploy an example application, we need to set it up in the udpdate service. You can read the detailed information in the [Getting Started]() guide, or quickly execute these commands:

```
$ updatectl update-app e96281a6-d1af-4bde-9a0a-97b76e56dc57 "Example Apache" "Sample application to test container updates"
$ updatectl new-group e96281a6-d1af-4bde-9a0a-97b76e56dc57 master 13DEFFFC-90C2-4B94-B3C2-1322BE8DC4E3 "Example Group"
$ updatectl new-package e96281a6-d1af-4bde-9a0a-97b76e56dc57 1.0 \
        --name foobar \
        --path coreos/example-apache:1.0 \
        --size 23 \
        --sha1sum fe7374bddde2ddf07f6bfcc728d115d14338964b \
        --sha256sum b602d630f0a081840d0ca8fc4d35810e42806642b3127bb702d65c3df227d0f5 \
        --signature ixi6Oebo \
        --metadata-size 190
$ updatectl update-channel e96281a6-d1af-4bde-9a0a-97b76e56dc57 master 1.0
```

In the web interface, you should see everything that you just created:

[screenshot]

## Deploy the Containers

Our application consists of a few main parts: a load balancer, a few Apache web servers, and both an updater and a service registration sidekick for each Apache container.

### Apache Unit

The main part of our application is the apache container, which simply prints a version when you browse to it. We'll use this to indicate when our containers update. In our unit, we've told docker to fetch `coreos/example-apache:1.0`. After an update is available, the updater will launch the new 2.0 containers.

Since we're running a distributed system, we don't want our frontends to run on the same machine. `X-Conflicts=apache-v1-*.service` will indicate to fleet that these services conflict and can't be run together. We'll run three of these, named `apache-1.service`, `apache-2.service` and `apache-3.service`.

```
[Unit]
Description=Apache v1.0
After=docker.service
Requires=docker.service

[Service]
ExecStart=/usr/bin/docker run -name apache -p 8080:80 coreos/example-apache:1.0 /usr/sbin/apache2ctl -D FOREGROUND
ExecStop=/usr/bin/docker stop apache

[X-Fleet]
# This unit will never live on the same box as another Apache unit
X-Conflicts=apache-*.service
```

### Service Registration Unit

The service registration unit will write the details of our Apache container into etcd in order for the load balancer to start serving traffic to it. Every 30 seconds, a heartbeat will refresh the entry in etcd. If the Apache container is stopped for any reason, this unit will also be stopped. We're going to use a simplified unit that doesn't even run a container, just a shell script:

```
[Unit]
Description=Announce Apache
# Binds this unit and an Apache unit together. When Apache is stopped, this unit will be stopped too.
BindsTo=%i.service

[Service]
EnvironmentFile=/etc/environment
ExecStart=/bin/sh -c "while true; do etcdctl set /services/apache/%i '{ \"host\": \"$COREOS_PUBLIC_IPV4\", \"port\": 8080 }' --ttl 60;sleep 30;done"
ExecStop=/usr/bin/etcdctl delete /services/apache/%i

[X-Fleet]
# This unit will always be colocated with the Apache unit
X-ConditionMachineOf=%i.service
```

### Update Unit

For each instance of our Apache container, we need to run an updater as a sidekick. Our simple updater is a bash script that can be found in the [example update repository]().

The `X-ConditionMachineOf` will indicate to fleet that this container *must* run on the same machine as our Apache container.

Since we're running multiple copies, each updater needs to reference a specific Apache instance. In order to use multiple instances of the same updater unit, we can use the variable `%i` to encode our reference into the filename. Naming our unit `updater@apache-1.service` will reference `apache-1.service`.

```
[Unit]
Description=Updater
After=docker.service
Requires=docker.service
# Binds this unit and an Apache unit together. When Apache is stopped, this unit will be stopped too.
BindsTo=%i.service

[Service]
ExecStart=/usr/bin/docker run -name updater coreos/example-updater /bin/sh /update.sh
ExecStop=/usr/bin/docker stop updater

[X-Fleet]
X-ConditionMachineOf=%i.service
```

### Load Balancer Container

vulcan

### Start The Units

Create three copies of the Apache unit and three copies of the updater unit, named appropriately. These unit files are also on [Github]().

```
$ ls
apache-1.service
apache-2.service
apache-3.service
updater@apache-1.service
updater@apache-2.service
updater@apache-3.service
```

Now start them with fleetctl. Be sure to only start the v1 Apache units:

```
fleetctl start apache-* updater@apache-*
```

## Verify Running Containers

There are three different ways we can verify that the units came up correctly. First, using fleet, we can `list-units`:

```
$ fleetctl list-units
INSERT ME
```

Second, open the load balancer in your browser and you should see version 1.0 being printed out. Refresh a few times in order to be routed to different containers.

By now the updaters should have reported back to the update service. Browse the the application and group you set up. Click on "View All Graphs" and change the time duration to hourly. You should see 3 Apache instances.

[screenshot]

### Troubleshooting the Updater

If you run into problems with the updater, the best way to diagnose issues is to read the journal. Using the `-follow` flag will continually display new log lines as they are generated:

```
fleetctl -follow updater@apache-1.service
```

For more journalctl information, check out the [Reading the System Log]({{site.url}}/docs/cluster-management/debugging/reading-the-system-log/) guide.


## Trigger Update

Now that we know everything is running, let's load a new package into the update service, configure the update behavior for the group, and update the channel to the new package version.

### Load the 2.0 Package

Create a new package pointing to the 2.0 tag of the container:

```
updatectl new-package e96281a6-d1af-4bde-9a0a-97b76e56dc57 2.0 \
	--name foobar \
	--path coreos/example-apache:2.0 \
	--size 23 \
	--sha1sum fe7374bddde2ddf07f6bfcc728d115d14338964b \
	--sha256sum b602d630f0a081840d0ca8fc4d35810e42806642b3127bb702d65c3df227d0f5 \
	--signature ixi6Oebo \
	--metadata-size 190
```

### Start the Update

In order to inspect our cluster mid-update, you can slow down the roll-out rate. Set the group to update 1 instance every 5 minutes. Once we promote the 2.0 package, we should see only one container update.

Now we'r ready to update. In the UI, modify the master channel to point to the 2.0 package that we loaded in. Since our updater is checking so frequently, we should see the new instance show up in the graph:

[screenshot]

We've successfully updated one of our containers! In your web browser, refresh the load balancer a few times. You should see a few requests coming from the updated container. Over the next 15 minutes, all of the containers should update.

## Further Reading

If you're interested in writing a more robust updater with customized logic for your applicaiton, check out the [protocol]() documentation.
