# sl-monitor [![Test](https://github.com/adikm/sl-monitor/actions/workflows/test.yml/badge.svg)](https://github.com/adikm/sl-monitor/actions/workflows/test.yml)


Storstockholms Lokaltrafik train monitor

This application serves API that allows to get information about Stockholm's region train traffic.
One can create and schedule notifications that inform about departures from chosen station and potential disturbances, if there are any.

### Prerequisites
Make sure you have Docker and SQLite3 installed locally on your machine


### Running
1. Prepare a SQLite database file, you can change the filename:
```shell
sqlite3 filename.db
```
2. Open [config.yml](config.yml) and configure environment variables as stated in the file.
   Optionally one can change the values in the config file as well.
3. Build the docker image:
```shell
make build
```

4. Run 
```shell 
make run DB_FILE filename.db
```

The application will start and should be accessible under the following link: ```http://localhost:4444```