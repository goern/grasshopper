FROM scratch

MAINTAINER Christoph Görn <goern@redhat.com>

LABEL RUN="docker run --tty --interactive --rm --privileged -v `pwd`:/greengrass -v /run:/run -v /:/host --net=host --name \${NAME} -e NAME=\${NAME} -e IMAGE=\${IMAGE} \${IMAGE}"

# add Grasshopper to the container image
ADD grasshopper /

# the entrypoint
ENTRYPOINT ["/grasshopper"]
