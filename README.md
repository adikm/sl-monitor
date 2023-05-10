# sl-monitor [![Test & deploy](https://github.com/adikm/sl-monitor/actions/workflows/github-ci.yml/badge.svg)](https://github.com/adikm/sl-monitor/actions/workflows/github-ci.yml)

Storstockholms Lokaltrafik train monitor

This application serves API that allows to get information about Stockholm's region train traffic.
One can create and schedule notifications that inform about departures from chosen station and potential disturbances,
if there are any.

### Prerequisites

The only prerequisite is Docker installed on your machine. Also you should get TrafikVerket API key.

### Architecture diagram

![Diagram](architecture.png)

### Terraforming Google Cloud Platform

Remember to login to GCP

```shell
gcloud auth application-default login
```

before running

```shell
terraform init
terraform apply
```

To synchronize with remote infrastructure, run:

```shell
terraform import module.project_services.google_project_service.project_services "oslogin.googleapis.com"
terraform import module.project_services.google_project_service.project_services "compute.googleapis.com"

terraform import google_compute_network.vpc_network slmonitor-network
terraform import google_compute_firewall.allow_ssh allow-ssh    
terraform import google_compute_address.static_ip slmonitor-instance
terraform import google_compute_instance.vm_instance slmonitor-instance
terraform import google_sql_database_instance.postgresql slmonitor-db1
terraform import google_sql_database.postgresql_db slmonitor/slmonitor-db1/slmonitor
terraform import google_sql_user.postgresql_user slmonitor/slmonitor-db1/postgres
terraform import google_redis_instance.slmonitor_cache slmonitor
```

### Running

1. Open [config.yml](config.yml) and configure environment variables as stated in the file.
   Optionally you can pass the variables directly to the _run_ command. If you wish to do so, skip this step.
2. Build local docker image: ```make build-deploy```
3. Run docker-compose

```shell
make run-deploy
```

optionally pass additional variables:

```shell 
env TRAFFIC_API_AUTH_KEY=value env MAIL_USERNAME=user env MAIL_PASSWORD=pass make run-deploy
```

All the services will start and the application should be accessible under the following
link: ```http://localhost:4444```