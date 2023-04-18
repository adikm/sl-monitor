# sl-monitor [![Test](https://github.com/adikm/sl-monitor/actions/workflows/test.yml/badge.svg)](https://github.com/adikm/sl-monitor/actions/workflows/test.yml)

Storstockholms Lokaltrafik train monitor

This application serves API that allows to get information about Stockholm's region train traffic.
One can create and schedule notifications that inform about departures from chosen station and potential disturbances,
if there are any.

### Prerequisites

The only prerequisite is Docker installed on your machine. Also you should get TrafikVerket API key.

### Architecture diagram

![Diagram](architecture.png)

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