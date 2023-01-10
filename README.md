# sl-monitor [![Test](https://github.com/adikm/sl-monitor/actions/workflows/test.yml/badge.svg)](https://github.com/adikm/sl-monitor/actions/workflows/test.yml)


Storstockholms Lokaltrafik train monitor

This application serves API that allows to get information about Stockholm's region train traffic.
One can create and schedule notifications that inform about departures from chosen station and potential disturbances, if there are any.

### Running
To be able to run it, open [config.yml](config.yml) and configure environment variables as stated in the file. 
Optionally one can change the values in the config file as well.

Then run
```go run .```

The server will start.