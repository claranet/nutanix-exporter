
BIN_NAME = nutanix-exporter
DOCKER_IMAGE_NAME ?= nutanix-exporter
GOPATH = $($pwd)

all: linux windows docker

linux: prepare
	$(eval GOOS=linux)
	$(eval GOARCH=amd64)
	go build -o ./bin/$(BIN_NAME)
	zip ./bin/$(BIN_NAME)-$(GOOS)-$(GOARCH).zip ./bin/$(BIN_NAME)

clean:
	@echo "Clean up"
	go clean
	rm -rf bin/

docker:
	@echo ">> Compile using docker container"
	@docker build -t "$(DOCKER_IMAGE_NAME)" .

windows: prepare
	$(eval GOOS=windows)
	$(eval GOARCH=amd64)
	go build -o ./bin/$(BIN_NAME).exe
	zip ./bin/$(BIN_NAME)-$(GOOS)-$(GOARCH).zip ./bin/$(BIN_NAME).exe
prepare:
	@echo "Create output directory ./bin/"
	mkdir -p bin/
	@echo "GO get dependencies"
	go get -d
	

.PHONY: all
