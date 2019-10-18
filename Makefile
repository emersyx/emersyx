emersyx: goget
	@go build -o emersyx ./main/*

.PHONY: goget
goget:
	@go get github.com/BurntSushi/toml
	@go get golang.org/x/lint/golint

.PHONY: test
test: emersyx
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l main/*.go api/*.go log/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
	@cd main; go test -v -conffile ../config.toml
