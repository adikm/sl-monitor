build:
	docker build --tag sl-monitor .
run:
	docker stop sl-monitor
	docker rm sl-monitor
	docker run --name sl-monitor --env TRAFFIC_API_AUTH_KEY=${TRAFFIC_API_AUTH_KEY} -p 4444:4444 sl-monitor
