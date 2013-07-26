NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

format:
	go fmt ./...

test:
	@echo "$(OK_COLOR)==> Testing tci...$(NO_COLOR)"
	@go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs -n1 go get -d
	go test ./...

.PHONY: format test
