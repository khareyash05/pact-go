export PACT_DIR = $(PWD)/pacts
GO          ?= go

build-order-service-consumer:
	$(GO) build ./...

run-order-service-consumer:
	$(GO) run cmd/main.go

run-order-service-provider:
	$(GO) run order-service-provider/cmd/main.go

consumer: export PACT_TEST := true


test-pact-consumer:
	$(GO) test ./...
