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
	cd $(PWD)/src && go build

run:
	cd $(PWD)/src && go build && mv sl-monitor $(PWD) && cd $(PWD) && ./sl-monitor

run-debug:
	cd $(PWD)/src && go build -gcflags "all=-N -l"  && mv sl-monitor $(PWD) && cd $(PWD) && dlv --listen=:4445 --headless=true --api-version=2 --accept-multiclient exec ./sl-monitor

test:
	cd $(PWD)/src && go test ./...