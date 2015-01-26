# Configure Machines to Use CoreUpdate

Configuring new or existing CoreOS machines to communicate with a [CoreUpdate](https://coreos.com/products/coreupdate) instance is a simple change to a configuration file.  Prior to making these changes ensure that you have the following information:

  - The server endpoint (in this example we will use `customer.update.core-os.net` as our example hostname)
  - The group identifier you wish to use. By default, the identifier for a group of machines is a UUID generated when the group was created. You may optionally specify a unique string which contains the characters `a-Z`, `0-9`, `-`, and `_`. This can be seen in practice with the default groups "alpha", "beta", and "stable" that map to those respective channels. The uniqueness of the string need only be in scope of the CoreUpdate deployment, not globally.

## New Machines

New servers can be configured to communicate with your CoreUpdate installation by using [cloud-config](https://coreos.com/docs/cluster-management/setup/cloudinit-cloud-config). Set the value of `server` to the custom address of your installation and `group` to the unique identifier of your application group:

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
