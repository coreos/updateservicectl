# API Client generation

The Update API client is automatically generated based on the API discovery
functionality provided by the endpoints library. This document describes how to
update `updateservicectl` to reflect changes in the API.

Most of the work is done by the Makefile. Before it can be run though, a couple
of things have to be set up. First, stand up an instance of CoreUpdate with the
API you want to generate the bindings for. Then, make sure you have the
`google-api-go-generator` tool installed. Get it with the following command - 

```
go get google.golang.org/api/google-api-go-generator
```

Once the environment is set up, simply run this make command - 

```
make api-gen
```

and the bindings should be updated. Check its work, commit all the binding
changes into a single, separate commit, and then update the necessary function
calls to use the new bindings (if necessary).
