# Configure Machines to Use CoreUpdate

Configuring new or existing CoreOS machines to communicate with a [CoreUpdate](https://coreos.com/products/coreupdate) instance is a simple change to a configuration file.

## New Machines

New servers can be configured to communicate with your CoreUpdate installation by using [cloud-config](https://coreos.com/docs/cluster-management/setup/cloudinit-cloud-config). Set the value of `server` to the custom address of your installation and `group` to the unique identifier of your application group.

For example, here is what "NYC Production" looks like in CoreUpdate:

![CoreUpdate Group](img/coreupdate-group.png)

Here's the cloud-config to use:

```
#cloud-config

coreos:
  update:
    group: 0a809ab1-c01c-4a6b-8ac8-6b17cb9bae09
    server: https://customer.update.core-os.net/v1/update/
```

More information can be found in the [cloud-config guide](http://coreos.com/docs/cluster-management/setup/cloudinit-cloud-config/#coreos).

## Existing Machines

To change the update of existing machines, edit `/etc/coreos/update.conf` with your favorite editor and provide the `SERVER=` and `GROUP=` values:

```
GROUP=0a809ab1-c01c-4a6b-8ac8-6b17cb9bae09
SERVER=https://customer.update.core-os.net/v1/update/
```

To apply the changes, run:

```
sudo systemctl restart update-engine
```

## Viewing Machines in CoreUpdate

Each machine should check in about 10 minutes after boot and roughly every hour after that. If you'd like to see it sooner, you can force an update check, which will skip any rate-limiting settings that are configured.

### Force Update in Background

```
$ update_engine_client -check_for_update
[0123/220706:INFO:update_engine_client.cc(245)] Initiating update check and install.
```

### Force Update in Foreground

If you want to see what's going on behind the scenes, you can watch the ouput in the foreground:

```
$ update_engine_client -update
[0123/222449:INFO:update_engine_client.cc(245)] Initiating update check and install.
[0123/222449:INFO:update_engine_client.cc(250)] Waiting for update to complete.
LAST_CHECKED_TIME=0
PROGRESS=0.000000
CURRENT_OP=UPDATE_STATUS_IDLE
NEW_VERSION=0.0.0.0
NEW_SIZE=0
[0123/222454:ERROR:update_engine_client.cc(189)] Update failed.
```

Be aware that the "failed update" means that there isn't a newer version to install.
