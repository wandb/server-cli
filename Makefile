PKG := github.com/wandb/server-cli
GOARCH := $(if $(GOARCH),$(GOARCH),"amd64")

.PHONY: deps

deps:
	go mod download
	go list -f '{{range .Imports}}{{.}} {{end}}' tools.go | GOOS=$(OS) xargs go install

check: deps
	go vet ./...
	staticcheck ./...
	ineffassign ./...

	(! go list -f {{.Dir}} ./... | xargs goimports -d -format-only -local=$(PKG) | read any) || { echo 'Formatting violations detected, run "make format"'; exit 1; }
	(! go list -f {{.Dir}} ./... | xargs golines --base-formatter=gofmt --max-len=120 -l | read any) || { echo 'Formatting violations detected, run "make format"'; exit 1; }

format: deps
	go list -f {{.Dir}} ./... | xargs goimports -format-only -w -local=$(PKG)
	go list -f {{.Dir}} ./... | xargs golines --base-formatter=gofmt --max-len=120 -w

test: deps
	go test ./...