export GO111MODULE=on

.PHONY: deps
deps:
	go get -u -d -v
	go mod tidy

.PHONY: devel-deps
devel-deps: deps
	sh -c '\
	tmpdir=$$(mktemp -d); \
	cd $$tmpdir; \
	go get ${u} \
	golang.org/x/lint/golint; \
	rm -rf $$tmpdir'

.PHONY: test
test: deps
	go test -v -cover ./...

.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...


.PHONY: build
build: deps
	go build -v

.PHONY: install
install: deps
	go install -v
