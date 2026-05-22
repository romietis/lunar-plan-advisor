IMAGE := lunar-plan-advisor
PORT  := 8080

.PHONY: build run up stop test e2e help

help:
	@echo "make build - build the Docker image ($(IMAGE))"
	@echo "make run   - run the container on port $(PORT)"
	@echo "make up    - build the image and run the container"
	@echo "make stop  - stop and remove the running container"
	@echo "make test  - run unit tests"
	@echo "make e2e   - run end-to-end (BDD) tests"

build:
	docker build -t $(IMAGE) .

run:
	docker run --rm --name $(IMAGE) -p $(PORT):$(PORT) $(IMAGE)

up: build run

stop:
	-docker stop $(IMAGE)

test:
	go test -v -shuffle=on -coverprofile=cover.out ./...

e2e:
	go test -v -tags=e2e ./internal/bdd/...
