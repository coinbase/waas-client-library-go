# All targets
.PHONY: all
all: clean protos mocks
	@go mod tidy

# Remove generated files
.PHONY: clean
clean:
	rm -rf ./gen

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
mocks: mocks/pools mocks/protocols mocks/blockchain mocks/mpc_keys mocks/mpc_wallets mocks/mpc_transactions

# Run the mockery script for the provided component
.PHONY: mocks/%
mocks/%:
	$(eval COMPONENT=$(@:mocks/%=%))
	@scripts/mockery.sh $(COMPONENT)

# Dependency protos.
.PHONY: .proto
.proto: .proto/googleapis .proto/openapiv2

# Google API protos.
.PHONY: .proto/googleapis
.proto/googleapis:
	@if [ ! -d .proto/googleapis ]; then \
		echo "Fetching Google API protos" ; \
		git clone -q --depth 1 --branch master https://github.com/googleapis/googleapis.git .proto/googleapis ; \
	fi

# OpenAPI v2 protos.
.PHONY: .proto/openapiv2
.proto/openapiv2:
	@if [ ! -d .proto/openapiv2 ]; then \
		echo "Fetching OpenAPI v2 protos" ; \
		git clone -q --depth 1 --branch main https://github.com/grpc-ecosystem/grpc-gateway.git .proto/openapiv2 ; \
	fi

# Lint API protos.
.PHONY: lint
lint: lint/pools lint/blockchain lint/protocols lint/mpc_keys lint/mpc_wallets lint/mpc_transactions

# Lint API protos.
.PHONY: lint/%
lint/%: .proto
	$(eval COMPONENT=$(@:lint/%=%))
	api-linter protos/services/coinbase/cloud/$(COMPONENT)/v1/$(COMPONENT).proto \
	-I .proto/googleapis -I .proto/openapiv2 -I protos/types --config api-linter.yml
