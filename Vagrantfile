# -*- mode: ruby -*-
# vi: set ft=ruby :

# The Docker registry from where we pull the OpenShift Docker image
DOCKER_REGISTRY="docker.io"

# The name of the OpenShift image.
ORIGIN_IMAGE_NAME="openshift/origin"

# The public IP address the VM created by Vagrant will get.
# You will use this IP address to connect to OpenShift web console.
PUBLIC_ADDRESS="10.1.2.2"
PUBLIC_HOST="adb.cluster.io"

# The directory where OpenShift will store the files.
# This should be "/var/lib/openshift" in case you're not using the :latest tag.
ORIGIN_DIR="/var/lib/origin"

Vagrant.configure(2) do |config|
  config.vm.box = 'centos/7'
  config.vm.provider "virtualbox" do |v, override|
    v.memory = 2048
    v.cpus   = 2
    v.customize ["modifyvm", :id, "--cpus", "2"]
    v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  end

  config.vm.provider "libvirt" do |v, override|
    v.driver = "kvm"
    v.memory = 2048
    v.cpus   = 2
  end

  config.vm.network "private_network", ip: "#{PUBLIC_ADDRESS}"


  # On OS X you can enable Landrush to expose OpenShift routes to the host
  # Install the Landrush plugin 'vagrant plugin install landrush'
  # and comment in the lines below
  # This won't work on Windows and on Linux the TLD ending on .local might cause
  # problems
  # config.vm.hostname = "#{PUBLIC_HOST}"
  # config.landrush.enabled = true
  # config.landrush.host_ip_address = "#{PUBLIC_ADDRESS}"
  # config.landrush.tld = "#{PUBLIC_HOST}"
  # config.landrush.guest_redirect_dns = false

  config.vm.provision "shell", inline: <<-SHELL
    yum install -y docker
    systemctl start docker
    docker inspect openshift/origin &>/dev/null && exit 0
    echo "[INFO] Pull the #{ORIGIN_IMAGE_NAME} Docker image ..."
    docker pull #{ORIGIN_IMAGE_NAME}
    docker tag #{DOCKER_REGISTRY}/#{ORIGIN_IMAGE_NAME} #{ORIGIN_IMAGE_NAME}
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    state=$(docker inspect -f "{{.State.Running}}" origin 2>/dev/null)
    [[ "${state}" == "true" ]] && exit 0
    if docker inspect origin &>/dev/null; then
      echo "[INFO] Removing previously started OpenShift server container ..."
      docker rm -f -v origin 2>/dev/null
    fi
    echo "[INFO] Start the OpenShift server ..."
    # Prepare directories for bind-mounting
    dirs=(openshift.local.volumes openshift.local.config openshift.local.etcd)
    for d in ${dirs[@]}; do
      mkdir -p #{ORIGIN_DIR}/${d} && chcon -Rt svirt_sandbox_file_t #{ORIGIN_DIR}/${d}
    done
    docker run -d --name "origin" --privileged --net=host --pid=host \
         -v /:/rootfs:ro \
         -v /var/run:/var/run:rw \
         -v /sys:/sys:ro \
         -v /var/lib/docker:/var/lib/docker:rw \
         -v #{ORIGIN_DIR}/openshift.local.volumes:#{ORIGIN_DIR}/openshift.local.volumes:z \
         -v #{ORIGIN_DIR}/openshift.local.config:#{ORIGIN_DIR}/openshift.local.config:z \
         -v #{ORIGIN_DIR}/openshift.local.etcd:#{ORIGIN_DIR}/openshift.local.etcd:z \
         #{ORIGIN_IMAGE_NAME} start \
          --master="https://#{PUBLIC_ADDRESS}:8443" \
          --etcd-dir="#{ORIGIN_DIR}/openshift.local.etcd" \
          --cors-allowed-origins=.*
    sleep 15 # Give OpenShift 15 seconds to start
    state=$(docker inspect -f "{{.State.Running}}" origin)
    if [[ "${state}" != "true" ]]; then
      >&2 echo "[ERROR] OpenShift failed to start:"
      docker logs origin
      exit 1
    fi
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    binaries=(oc oadm)
    for n in ${binaries[@]}; do
      [ -f /usr/bin/${n} ] && continue
      echo "[INFO] Copy the OpenShift '${n}' binary to host /usr/bin/${n}..."
      docker run --rm --entrypoint=/bin/cat #{ORIGIN_IMAGE_NAME} /usr/bin/${n} > /usr/bin/${n}
      chmod +x /usr/bin/${n}
    done
    echo "export OPENSHIFT_DIR=#{ORIGIN_DIR}/openshift.local.config/master" > /etc/profile.d/openshift.sh
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    export KUBECONFIG=${OPENSHIFT_DIR}/admin.kubeconfig
    chmod go+r ${KUBECONFIG}
    # Create Docker Registry
    if [ ! -f #{ORIGIN_DIR}/configured.registry ]; then
      echo "[INFO] Configure Docker Registry ..."
      oadm registry --create --credentials=${OPENSHIFT_DIR}/openshift-registry.kubeconfig
      touch #{ORIGIN_DIR}/configured.registry
    fi
    # For router, we have to create service account first and then use it for
    # router creation.
    if [ ! -f #{ORIGIN_DIR}/configured.router ]; then
      echo "[INFO] Configure HAProxy router ..."
      echo '{"kind":"ServiceAccount","apiVersion":"v1","metadata":{"name":"router"}}' \
        | oc create -f -
      oc get scc privileged -o json \
        | sed '/\"users\"/a \"system:serviceaccount:default:router\",'  \
        | oc replace scc privileged -f -
      oadm router --create --credentials=${OPENSHIFT_DIR}/openshift-router.kubeconfig \
        --service-account=router
      touch #{ORIGIN_DIR}/configured.router
    fi
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    export KUBECONFIG=${OPENSHIFT_DIR}/admin.kubeconfig
    chmod go+r ${KUBECONFIG}
    if [ ! -f ${ORIGIN_DIR}/configured.templates ]; then
      echo "[INFO] Installing OpenShift templates"

      # TODO - These list must be verified and completed for a official release
      # Currently templates are sources from three main repositories
      # - openshift/origin
      # - openshift/nodejs-ex
      # - jboss-openshift/application-templates
      ose_tag=ose-v1.1.0
      template_list=(
        # Image streams
        https://raw.githubusercontent.com/openshift/origin/master/examples/image-streams/image-streams-rhel7.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/jboss-image-streams.json
        # DB templates
        https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mongodb-ephemeral-template.json
        https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mongodb-persistent-template.json
        https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mysql-ephemeral-template.json
        https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/mysql-persistent-template.json
        https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/postgresql-ephemeral-template.json
        https://raw.githubusercontent.com/openshift/origin/master/examples/db-templates/postgresql-persistent-template.json
        # Jenkins
        https://raw.githubusercontent.com/openshift/origin/master/examples/jenkins/jenkins-ephemeral-template.json
        https://raw.githubusercontent.com/openshift/origin/master/examples/jenkins/jenkins-persistent-template.json
        # Node.js
        https://raw.githubusercontent.com/openshift/nodejs-ex/master/openshift/templates/nodejs-mongodb.json
        https://raw.githubusercontent.com/openshift/nodejs-ex/master/openshift/templates/nodejs.json
        # EAP
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-amq-persistent-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-amq-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-basic-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-https-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-mongodb-persistent-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-mongodb-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-mysql-persistent-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-mysql-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-postgresql-persistent-s2i.json
        https://raw.githubusercontent.com/jboss-openshift/application-templates/${ose_tag}/eap/eap64-postgresql-s2i.json
      )

      for template in ${template_list[@]}; do
        echo "[INFO] Importing template ${template}"
        oc create -f $template -n openshift >/dev/null
      done
      touch ${ORIGIN_DIR}/configured.templates
    fi
  SHELL

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    echo "[INFO] Create 'test-admin' user and 'test' project ..."
    if [ ! -f #{ORIGIN_DIR}/configured.user ]; then
      oadm policy add-role-to-user view test-admin --config=${OPENSHIFT_DIR}/admin.kubeconfig
      oc login https://#{PUBLIC_ADDRESS}:8443 -u test-admin -p test \
        --certificate-authority=${OPENSHIFT_DIR}/ca.crt &>/dev/null
      oc new-project test --display-name="OpenShift 3 Sample" \
        --description="This is an example project to demonstrate OpenShift v3" &>/dev/null
      sudo touch #{ORIGIN_DIR}/configured.user
    fi
    echo
    echo "You can now access OpenShift console on: https://#{PUBLIC_ADDRESS}:8443/console"
    echo
    echo "To use OpenShift CLI, run:"
    echo "$ vagrant ssh"
    echo "$ oc status"
    echo
    echo "To become a cluster-admin, add '--config' to oc commands:"
    echo "$ vagrant ssh"
    echo "$ oc status --config=${OPENSHIFT_DIR}/admin.kubeconfig"
    echo
    echo "To browse the OpenShift API documentation, follow this link:"
    echo "http://openshift3swagger-claytondev.rhcloud.com"
    echo
    echo "Then enter this URL:"
    echo https://#{PUBLIC_ADDRESS}:8443/swaggerapi/oapi/v1
    echo "."
  SHELL

end
