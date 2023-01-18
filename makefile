build:
	docker build --tag sl-monitor .
run:
	docker run -d --name sl-monitor -e TRAFFIC_API_AUTH_KEY=${TRAFFIC_API_AUTH_KEY} -e EMAIL_USERNAME=${EMAIL_USERNAME} -e EMAIL_PASSWORD=${EMAIL_PASSWORD} --volume ${PWD}/${DB_FILE}:/database.db -p 4444:4444 sl-monitor
stop:
	docker stop sl-monitor
	docker rm sl-monitor