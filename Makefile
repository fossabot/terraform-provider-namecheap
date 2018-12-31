default: build

build:
	go build -o bin/terraform-provider-namecheap .

test:
	go test ./...

test-race:
	go test -race ./...

test-cover:
	go test -cover -coverprofile=coverage.txt -covermode=atomic ./...

fmtcheck:
	diff -u <(echo -n) <(gofmt -d -s .)

modcheck:
	go mod tidy && git diff --exit-code; code=$?; git checkout -- .; (exit $code)

lint:
	golangci-lint run
