.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=5 go build

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: inst-golangci-lint
inst-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
