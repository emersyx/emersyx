emersyx:
	@go build -o emersyx ./cmd/emersyx/

.PHONY: test
test: emersyx
	@echo "Running the tests with gofmt."
	@test -z "$(shell gofmt -s -l	\
		cmd/emersyx/*.go			\
		pkg/api/*.go				\
		pkg/api/irc/*.go			\
		pkg/api/telegram/*.go		\
		pkg/log/*.go				\
	)"
	@echo "Running the tests with go vet."
	@go vet ./...
	@echo "Running the tests with golint."
	@golint -set_exit_status $(shell go list ./...)
	@cd cmd/emersyx; go test -v -conffile ../../config/config-template.toml
