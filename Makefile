.PHONY: tidy
# Upgrade packages
tidy:
	go mod tidy -v

.PHONY: build
# Upgrade packages
build:
	go build ./...

.PHONY: upgrade
# Upgrade packages
upgrade:
	go get -u -v -t ./...
	$(MAKE) tidy

.PHONY: upgrade2
# Upgrade packages by using go-mod-upgrade
upgrade2:
	goup -v || go-mod-upgrade -v

up-build: upgrade build

.PHONY: lint
# Run golangci-lint
lint:
	golangci-lint run --config .golangci.yaml ./...

.PHONY: fmt
# Run formatters
fmt:
	golangci-lint fmt --diff

.PHONY: config-verify
# Validate golangci-lint configuration
config-verify:
	golangci-lint config verify