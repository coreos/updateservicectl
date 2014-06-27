FROM busybox
# Note that the updatectl in the current directory MUST be built with GOOS=linux
ADD updatectl /usr/bin/updatectl
ENTRYPOINT ["/usr/bin/updatectl"]
