BINARY_NAME=app

build:
	echo "Building..."
	go mod download
	go mod verify
	go build -o ${BINARY_NAME} cmd/main.go
	echo "Finished..."

clean:
	go clean
	rm ${BINARY_NAME}

lint:
	go vet ./...
	staticcheck ./...
	golint ./...

env: 
	docker-compose up -d

docker-build:
	docker build -t app . -f Dockerfile.dev

docker-run:
	make env
	docker run --net=host --rm -p 8013:8013 app:latest

run:
	make build
	./${BINARY_NAME}

test:
	go test -v ./...
	go test ./... -coverprofile=coverage.out
