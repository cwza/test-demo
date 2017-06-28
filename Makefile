TEST_PACKAGES := $(shell go list ./... | grep -vE '(vendor|cmd)')

test-unit:
	go test -short ${TEST_PACKAGES}

test:
	go test ${TEST_PACKAGES}

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main ./cmd/main.go
