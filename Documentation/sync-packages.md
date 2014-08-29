# Synchronizing CoreOS Packages

If you have an on-premise deployment of CoreUpdate behind a firewall, you can use the `updateservicectl` tool to synchronize packages and payloads.
*See the [upstream management docs](https://github.com/coreos/updateservicectl/blob/master/Documentation/client.md#upstream-management) if you wish to connect your instance to the internet and enable automatic synchronization.*

Non-firewalled instances of CoreUpdate do not require this step because they will synchronize automatically over the internet.

## Download Packages

From a machine with internet access, download the packages from the public CoreUpdate instance:  
```bash
updateservicectl --server=https://public.update.core-os.net package download --dir=./coreos-packages
```

## Upload Packages

There are two ways to upload pacakges depending on your storage configuration.

### Payloads Stored with CoreUpdate

This requires the `ENABLE_PACKAGE_UPLOADS` and `STATIC_PACKAGES_DIR` options to have been set during initial configuration.

First upload the package meta-data:  
```bash
updateservicectl --server=http://your-server.com package create bulk --dir=./coreos-packages
```

Next upload the actual package payloads:  
```bash
updateservicectl --server=http://your-server.com package upload bulk --dir=./coreos-packages
```

### Payloads Stored with Custom Storage

First upload the package meta-data:  
```bash
updateservicectl --server=http://your-server.com package create bulk --dir=./coreos-packages --base-url=http://custom-file-server
```

Note that the `--base-url` flag must be specified in order to rewrite the download location of the packages.
This should contain the full URL to the package excluding the filename.
Do not alter the base filenames.

Next you will need to copy the files you downloaded previously to your custom file store.
(It is not necessary to copy the json files, only the actual payload files are required)
