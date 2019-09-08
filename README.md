# Test tasks for S01

## Firts task

Is superlate and this is last task for me, so sorry for bash script only, but should work.

Copy nagios_es_plugin.cfg file to /etc/nagios-plugins/config/ dir, and add to main nagios cfg.

Next copy kafka_check to /usr/lib/nagios/plugins/ and 

`chmod u+x /usr/lib/nagios/plugins/elastic_check.sh`

## Second task

Copy nagios_kafka_plugin.cfg file to /etc/nagios-plugins/config/ dir, and add to main nagios.cfg

For compliation you will have install go lang in your os:  <https://golang.org/doc/install>

After that just simply run in this directory:

`make`

Next copy nagios/kafka/kafka_check to /usr/lib/nagios/plugins/ and:

`chmod u+x /usr/lib/nagios/plugins/kafka_check`

## Third task

### Dev

Setup kubectl for dev cluster

`gcloud container clusters get-credentials s01 --zone europe-west1-b --project dev-242510`

Setup secrets and namespace:

`kubectl create secret generic mysqlrootpassword --from-literal=password="password"`

`kubectl create secret generic mysqluserpassword --from-literal=password="password"`

`kubectl create namespace springboot`

Create mysql pv and mysql database depolyment:

`kubect create -f k8s/dev/mysql-pv.yaml`

`kubect create -f k8s/dev/mysql.yaml`

Create deployment for springboot deployment with external loadbalancer:

`kubect create -f k8s/dev/springboot.yaml`
In this case (for instnace for minicube) you have to provide external loadbalancer like ie.: HAPROXY

OR

Create deployment for springboot deployment with loadbalancer by cloud vendor:

`kubect create -f k8s/dev/springboot-lb-version.yaml`

In this case cloud vendor will provide loadbalancer and in some cases additional ingress will be needed.

### QA

Setup kubectl for dev cluster

`gcloud container clusters get-credentials s01 --zone europe-west1-b --project qa-242510`

Setup secrets and namespace:

`kubectl create secret generic mysqlrootpassword --from-literal=password="password"`

`kubectl create secret generic mysqluserpassword --from-literal=password="password"`

`kubectl create namespace springboot`

Create mysql pv and mysql database depolyment:

`kubect create -f k8s/qa/mysql-pv.yaml`

`kubect create -f k8s/qa/mysql.yaml`

Create deployment for springboot deployment with external loadbalancer:

`kubect create -f k8s/qa/springboot.yaml`
In this case (for instnace for minicube) you have to provide external loadbalancer like ie.: HAPROXY

OR

Create deployment for springboot deployment with loadbalancer by cloud vendor:

`kubect create -f k8s/qa/springboot-lb-version.yaml`

In this case cloud vendor will provide loadbalancer and in some cases additional ingress will be needed.

### Prod

Setup kubectl for dev cluster

`gcloud container clusters get-credentials s01 --zone europe-west1-b --project prod-242510`

Setup secrets and namespace:

`kubectl create secret generic mysqlrootpassword --from-literal=password="password"`

`kubectl create secret generic mysqluserpassword --from-literal=password="password"`

`kubectl create namespace springboot`

Create mysql pv and mysql database depolyment:

`kubect create -f k8s/prod/mysql-pv.yaml`

`kubect create -f k8s/prod/mysql.yaml`

Create deployment for springboot deployment with external loadbalancer:

`kubect create -f k8s/prod/springboot.yaml`
In this case (for instnace for minicube) you have to provide external loadbalancer like ie.: HAPROXY

OR

Create deployment for springboot deployment with loadbalancer by cloud vendor:

`kubect create -f k8s/prod/springboot-lb-version.yaml`

In this case cloud vendor will provide loadbalancer and in some cases additional ingress will be needed.

#### HA

I'm too tired for create second mysql deployment and setup replication between them or setup Galera cluster from mariadb project.

#### CI

Example of CI/CD for gitlab is in gitlab-ci.yaml file

#### Moniotoring

I would use sidecar containter in pod with prometheus node_exporter, and send data to grafana (cpu, memory, storage space, nr open files/files descriptors, etc....).

Also with developers is possible to add export prometheus directly from application.
