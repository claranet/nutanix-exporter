
BIN_NAME = nutanix-exporter
DOCKER_IMAGE_NAME ?= nutanix-exporter
GOPATH = $($pwd)

all: build clean

build:
	@echo "Create output directory ./bin/"
	mkdir -p bin/
	@echo "GO get dependencies"
	go get -d
	@echo "Build ..."
	go build -o ./bin/$(BIN_NAME)

clean:
	@echo "Clean up"
	@go clean

docker:
	@echo ">> Compile using docker container"
	@docker build -t "$(DOCKER_IMAGE_NAME)" .

deploy:
	@echo ">> Deploy docker container"
	@docker push ${DOCKER_IMAGE_NAME}

.PHONY: all
