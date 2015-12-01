FROM scratch

MAINTAINER Christoph Görn <goern@redhat.com>

LABEL Component="grasshopper" \
      Name="goern/grasshopper-0-generic" \
      Version="0.1.0" \
      Release="1"

LABEL io.k8s.description="This is Grasshopper, it will make your Nulecule GO!" \
      io.k8s.display-name="Grasshopper 0.1.0" \
      io.openshift.tags="grasshopper,nulecule,atomicapp"

LABEL RUN="docker run --tty --interactive --rm --privileged -v `pwd`:/greengrass -v /run:/run -v /:/host --net=host --name \${NAME} -e NAME=\${NAME} -e IMAGE=\${IMAGE} \${IMAGE}"

# add Grasshopper to the container image
ADD grasshopper /

# the entrypoint
ENTRYPOINT ["/grasshopper"]
