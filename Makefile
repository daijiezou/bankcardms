VERSION ?= $(shell git rev-parse --abbrev-ref HEAD)
MAKEFLAGS += --jobs all
LDFLAGS:=--ldflags "\
		-X 'main.Version=$(VERSION)' \
		-X 'main.BuildNo=$(shell git rev-parse HEAD)' \
		-X 'main.BuildTime=$(shell date +'%Y-%m-%d %H:%M:%S')'"

define push_image
        docker tag ${1} 10.10.10.149:80/${1}
        docker push 10.10.10.149:80/${1}
        docker rmi 10.10.10.149:80/${1}

endef


default: tidy local
tidy:
	go mod tidy -compat=1.18
test:
	go test ./...

build: local

local: tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -o bank-card-ms ${LDFLAGS} main.go

run:
	go run main.go  -env release -config ./config/config.yml

.PHONY: image
image:
	docker build -f Dockerfile -t bank-card-ms:dev-manual .
	$(call push_image,bank-card-ms:dev-manual)
	rm bank-card-ms

