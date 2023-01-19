# docker operations
build-deploy:
	docker build --tag sl-monitor .

run-deploy:
	docker run -d --name sl-monitor -e TRAFFIC_API_AUTH_KEY=${TRAFFIC_API_AUTH_KEY} -e EMAIL_USERNAME=${EMAIL_USERNAME} -e EMAIL_PASSWORD=${EMAIL_PASSWORD} --volume ${PWD}/${DB_FILE}:/database.db -p 4444:4444 sl-monitor

stop-deploy:
	docker stop sl-monitor
	docker rm sl-monitor

# local operations
build:
	go build -v

run:
	go run .

test:
	go test ./..