FROM scratch

MAINTAINER goern@redhat.com

# add all of Atomic App's files to the container image
ADD grasshopper /

# the entrypoint
ENTRYPOINT ["/grasshopper"]
