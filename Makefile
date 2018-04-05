emersyx: goget
	@go build -o emersyx ./core/* ./router/*

.PHONY: goget
goget:
	@go get emersyx.net/emersyx_log/emlog
	@go get github.com/BurntSushi/toml
	@go get github.com/golang/lint/golint

.PHONY: test
test: emersyx
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l core/*.go api/*.go router/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
	@cd core; go test -v -conffile ../config.toml
