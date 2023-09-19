#GOBIN?=.
GOBIN?=/src/bin
GO111MODULE?=on
PROJECT_NAME?=cocas
PROJECT_PATH?=coCAS

setting-prepare:
	go mod tidy
	go mod vendor

install-devtools: setting-prepare
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest
	go install github.com/segmentio/golines@latest
	which swagger || (GO111MODULE=on go install github.com/go-swagger/go-swagger/cmd/swagger@latest)

run-test:
	go clean -testcache
	go test ./... -covermode=count -coverprofile ./coverage.out

run-fmt:
	go fmt
	$(HOME)/go/bin/yamlfmt test/
	$(HOME)/go/bin/golines . -w

start-app: setting-prepare
	go run .

build-app: setting-prepare
	go build -mod=vendor -o $(GOBIN)/$(PROJECT_PATH)

install-app: build-app
	go install

docker-build:
	docker build . -f docker/Dockerfile --tag $(PROJECT_NAME):latest

docker-start:
	docker rm -f $(PROJECT_NAME) 2> /dev/null
	docker run \
		--name $(PROJECT_NAME) \
		-v $(PWD)/test:/etc/$(PROJECT_PATH)/conf \
		-p 8181:8181 \
		$(PROJECT_NAME):latest \
		-conf=/etc/$(PROJECT_PATH)/conf/conf.yaml
