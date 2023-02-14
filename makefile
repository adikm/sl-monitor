# docker operations
build-deploy:
	docker build --tag sl-monitor .

run-deploy:
	docker run -d --name sl-monitor -e TRAFFIC_API_AUTH_KEY=${TRAFFIC_API_AUTH_KEY} -e MAIL_USERNAME=${MAIL_USERNAME} -e MAIL_PASSWORD=${MAIL_PASSWORD} --volume ${PWD}/${DB_FILE}:/database.db -p 4444:4444 sl-monitor

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