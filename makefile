PACKAGE_NAME=http-spec

all: test build
build: build-local build-linux
	echo "building..."
build-local:
	echo "building local..."
	go build -o $(PACKAGE_NAME) -i -ldflags '-s -w'
build-linux:
	echo "building linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(PACKAGE_NAME) -i -ldflags '-s -w'
build-docker: build-linux
	echo "building docker..."
	docker build --tag tmornini/$(PACKAGE_NAME) .
test:
	echo "testing..."
	go test -v ./...
test-docker: clean-docker build-docker
	echo "testing docker..."
	docker build -t tmornini/validate-http-spec -f examples/example-Dockerfile .
	docker run --rm tmornini/validate-http-spec
clean:
	echo "cleaning..."
	go clean
	rm -f $(PACKAGE_NAME)
clean-docker:
	echo "cleaning docker..."
	docker 2>/dev/null 1>&2 rmi tmornini/validate-http-spec || true
	docker 2>/dev/null 1>&2 rmi tmornini/http-spec || true
run: build
	echo "running..."
	./$(PACKAGE_NAME)
