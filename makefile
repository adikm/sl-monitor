# docker operations
build-deploy:
	docker build --tag sl-monitor .

run-deploy:
	env DB_HOST="host.docker.internal" docker-compose -f docker-compose.yml up -d

stop-deploy:
	docker-compose stop

# local operations
build:
	cd $(PWD)/src && go build

run:
	cd $(PWD)/src && go build && mv sl-monitor $(PWD) && cd $(PWD) && ./sl-monitor

run-debug:
	cd $(PWD)/src && go build -gcflags "all=-N -l"  && mv sl-monitor $(PWD) && cd $(PWD) && dlv --listen=:4445 --headless=true --api-version=2 --accept-multiclient exec ./sl-monitor

test:
	cd $(PWD)/src && go test ./...