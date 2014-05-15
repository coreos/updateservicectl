# Omaha

The Omaha protocol is the specification that the update service uses to communicate with updaters running in a CoreOS cluster. The protocol is a fairly simple &mdash; it specifies sending HTTP POSTs with XML data bodies for various events that happen during the execution of an update.

## Update Request

The update request sends machine metadata and a list of applications that it is responsible for. In most cases, each updater is responsible for a single package. Here's what a typical request looks like:

```
<?xml version="1.0" encoding="UTF-8"?>
<request protocol="3.0">
 <app appid="e96281a6-d1af-4bde-9a0a-97b76e56dc57" version="1.0.0" track="beta" bootid="{fake-client-018}">
  <event eventtype="3" eventresult="2"></event>
 </app>
</request>
```

### Application Section

The app section is where the action happens. You can submit multiple applications or application instances in one request, but this isn't standard.

| Parameter | Description |
|-----------|-------------|
| appid     | Matches the id of the group that that this instance belongs to in the update service. |
| version   | The current semantic version number of the application code. |
| track     | The channel that the application is requesting. |
| bootid    | The unique identifier assigned to this instance. |

## Already Up to Date

If the application instance is already running the latest version, the response will be short:

```
<?xml version="1.0" encoding="UTF-8"?>
<response protocol="3.0" server="update.core-os.net">
 <daystart elapsed_seconds="0"></daystart>
 <app appid="e96281a6-d1af-4bde-9a0a-97b76e56dc57" status="ok">
  <updatecheck status="noupdate"></updatecheck>
 </app>
</response>
```

As you can see, the response indicated that no update was required for the provided group id and version.

## Update Required

If the application is not up to date, the response returned contains all of the information needed to execute the update:

```
<?xml version="1.0" encoding="UTF-8"?>
<response protocol="3.0" server="update.core-os.net">
 <daystart elapsed_seconds="0"></daystart>
 <app appid="e96281a6-d1af-4bde-9a0a-97b76e56dc57" status="ok">
  <updatecheck status="ok">
   <urls>
    <url codebase="http://index.example.com/webapp:1.0.2"></url>
   </urls>
   <manifest version="1.0.2">
    <packages>
     <package hash="fe7374bddde2ddf07f6bfcc728d115d14338964b" name="update.gz" size="23" required="false"></package>
    </packages>
    <actions>
     <action event="postinstall" sha256="b602d630f0a081840d0ca8fc4d35810e42806642b3127bb702d65c3df227d0f5" needsadmin="false" IsDelta="false" DisablePayloadBackoff="true" MetadataSignatureRsa="ixi6Oebo" MetadataSize="190"></action>
    </actions>
   </manifest>
  </updatecheck>
 </app>
</response>
```

The most important parts of the response are the `codebase`, which points to the location of the package, and the `sha256` which should be checked to make sure the package hasn't been tampered with.

## Report Progress, Errors and Completion

[not sure how to capture these...]
must report progress every X seconds
must report all errors
must report success

## Further Reading

Check out the [example container update workflow]() to see the protocol in action. You can read more about the [Omaha tech specs](https://code.google.com/p/omaha/wiki/ServerProtocol) or visit the [project homepage](https://code.google.com/p/omaha/).
