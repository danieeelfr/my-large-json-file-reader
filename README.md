# my-large-json-file-reader
This application is a tech challenge where I created this large-JSON-file-reader

This is a simple go application that can read a large json file and store each record in a redis memory database.

### Preparing the environment
```bash
# Use docker-compose to set up the environment locally
docker-compose up -d
```

## Building  
### with docker

```bash
# Build a docker image using DockerFile
docker build -t app . -f Dockerfile.dev
```

### without docker
```bash
# Build a binary file
go build -o app cmd/main.go
```

## Running  
### with docker

```bash
# Running the docker image created previously
docker run --net=host --rm -p 8013:8013 app:latest
```

### without docker
```bash
# Running the binary file built previously
./app
```
## Make commands
```bash
# building the binary file
make build  
# cleaning the binary file
make clean
# building the application docker image
make docker-build 
# building and running the application docker image
make docker-run 
# preparing the local environment
make env 
# linting the application code
make lint 
# running the application
make run 
# executing unit tests
make test
```

## Environment variables
```bash
# JSON FILE PATH is the path of the json file
JSON_FILE_PATH=

# REDIS connection string data
REDIS_HOST=
REDIS_PORT=
REDIS_PASSWORD=
REDIS_DB=

```
