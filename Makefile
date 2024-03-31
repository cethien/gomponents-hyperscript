.PHONY: format
format:
	@go fmt ./... && \
	go run golang.org/x/tools/cmd/goimports@latest -w .

.PHONY: lint
lint:
	@go vet ./... && \
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks all ./...

.PHONY: cover
cover:
	@go tool cover -html=cover.out

.PHONY: test
test:
	@go test -coverprofile=cover.out -shuffle on ./...
