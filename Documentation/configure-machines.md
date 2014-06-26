# Configure Machines to Use CoreUpdate

Configuring new or existing CoreOS machines to communicate with a [CoreUpdate](https://coreos.com/products/coreupdate) instance is a simple change to a configuration file.

## New Machines

New servers can be configured to communicate with your CoreUpdate installation by using [cloud-config](https://coreos.com/docs/cluster-management/setup/cloudinit-cloud-config). Set the value of `SERVER=` to the custom address of your installation and `GROUP=` to the unique identifier of your application group:

```
#cloud-config

write_files:
  - path: /etc/coreos/update.conf
    permissions: 0644
    content: |
      GROUP=0a809ab1-c01c-4a6b-8ac8-6b17cb9bae09
      SERVER=https://customer.update.core-os.net
```

## Existing Machines

To change the update of existing machines, edit `/etc/coreos/update.conf` with your favorite editor and provide the `SERVER=` and `GROUP=` values:

```
GROUP=0a809ab1-c01c-4a6b-8ac8-6b17cb9bae09
SERVER=https://customer.update.core-os.net
```

To apply the changes, run:

```
sudo systemctl restart update-engine
```