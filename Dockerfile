FROM scratch

MAINTAINER goern@redhat.com

# add Grasshopper to the container image
ADD grasshopper /

# the entrypoint
ENTRYPOINT ["/grasshopper"]
