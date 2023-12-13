EXECUTABLE := staterecovery
DEPLOY_ACCOUNT := degate
DEPLOY_IMAGE := $(EXECUTABLE)

.PHONY: clean
all: clean build_recover_cli copy_config

clean:
	rm -rf output/

build_recover_cli:
	mkdir -p output/
	go build -o output/staterecovery ./cmd/main.go

copy_config:
	mkdir -p output/conf
	cp -avR cmd/conf/* output/conf/

docker_build: copy_config
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags "$(EXTLDFLAGS)-s -w $(LDFLAGS)" -o output/$(EXECUTABLE) ./cmd/main.go

docker_image: docker_build
	docker build -t $(DEPLOY_ACCOUNT)/$(DEPLOY_IMAGE) -f docker/Dockerfile .

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o release/linux/amd64/$(EXECUTABLE) ./cmd/main.go

build_linux_arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o release/linux/arm64/$(EXECUTABLE) ./cmd/main.go

build_linux_arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -o release/linux/arm/$(EXECUTABLE) ./cmd/main.go

build_mac_intel:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o release/mac/intel/$(EXECUTABLE) ./cmd/main.go

build_mac_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -o release/mac/arm64/$(EXECUTABLE) ./cmd/main.go

build_windows_64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o release/windows/intel/$(EXECUTABLE).exe ./cmd/main.go