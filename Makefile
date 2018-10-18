NAME="ffp-updater-service"

lint:
	go get -u golang.org/x/lint/golint
	golint ./...

test-unit:
	go test -race -v -cover ./...

test: test-unit

build:
	go build -race -o bin/${NAME}

run-local:
	go build -race -o bin/${NAME}
	./bin/${NAME}

run-docker:
	docker-compose up --build 
