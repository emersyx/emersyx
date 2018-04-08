emersyx: goget
	@go build -o emersyx ./core/*

.PHONY: goget
goget:
	@go get github.com/BurntSushi/toml
	@go get github.com/golang/lint/golint

.PHONY: test
test: emersyx
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l core/*.go api/*.go router/*.go log/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
	@cd core; go test -v -conffile ../config.toml
