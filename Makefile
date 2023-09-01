export GO111MODULE=on
export GOOS=linux
export GOARCH=amd64

GIT_COMMIT := $(shell git rev-parse --short=8 HEAD)
VERSION ?= ${GIT_COMMIT}_`date '+%Y%m%d%H%M'`

GO_LDFLAGS += -X 'common.Version=${VERSION}'

version:
	@echo ${VERSION}

influxdb-test:
	go build  -ldflags "$(GO_LDFLAGS)" -o ./admin/main ./admin/main.go
	docker build  -t influxdb-test:${VERSION} --no-cache .
	docker tag influxdb-test:${VERSION} csighub.tencentyun.com/nosql_test/influxdb-test:${VERSION}
	docker push csighub.tencentyun.com/nosql_test/influxdb-test:${VERSION}
	rm ./admin/main

tidy:
	go mod verify
	go mod tidy
	@if ! git diff --quiet go.mod go.sum; then \
		echo "please run go mod tidy and check in changes, you might have to use the same version of Go as the CI"; \
		exit 1; \
	fi

default: help
help:
	@echo 'Targets:'
	@echo '  influxdb-test       - compile  influxdb-test image'
	@echo '  tidy      		   - tidy go modules'