export PACT_DIR = $(PWD)/pacts
GO          ?= go

provider: export PACT_TEST := true

run-order-service-provider:
	$(GO) run order-service-provider/cmd/main.go

test-pact-provider:
	@echo "Running tests..."
	$(GO) test ./...