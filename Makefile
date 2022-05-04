BINARY_NAME=main
GO_TEST_COMMAND=go test ./... -v
GO_TEST_COVERAGE_TARGET_FILE=coverage/coverage.out

clean:
	go clean
	rm -rf bin

install:
	go mod download

compile:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} .
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin .
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows .

build: clean install compile

run:
	./bin/${BINARY_NAME}

dev:
	go run .

test:
	${GO_TEST_COMMAND}

test_with_coverage:
	${GO_TEST_COMMAND} -coverprofile=${GO_TEST_COVERAGE_TARGET_FILE}

pre_test_coverage:
	rm -rf coverage
	mkdir coverage

post_test_coverage:
	go tool cover -html=${GO_TEST_COVERAGE_TARGET_FILE}

test_coverage: pre_test_coverage test_with_coverage post_test_coverage
