# All targets
.PHONY: all
all: protos mocks
	@go mod tidy

# Run all proto plugins
.PHONY: protos
protos: protos/types protos/services

# Generate type protos
.PHONY: protos/types
protos/types:
	@cd protos/types && buf generate

# Generate service protos
.PHONY: protos/services
protos/services:
	@cd protos/services && buf generate

# Generate Server and Client mocks
.PHONY: mocks
mocks: mocks/pools mocks/protocols mocks/blockchain

# Run the mockery script for the provided component
.PHONY: mocks/%
mocks/%:
	$(eval COMPONENT=$(@:mocks/%=%))
	@scripts/mockery.sh $(COMPONENT)
