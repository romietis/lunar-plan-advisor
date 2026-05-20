IMAGE := lunar-plan-advisor
PORT  := 8080

.PHONY: build run stop test e2e help

help:
	@echo "make build - build the Docker image ($(IMAGE))"
	@echo "make run   - run the container on port $(PORT) (detached)"
	@echo "make stop  - stop and remove the running container"
	@echo "make test  - run unit tests"
	@echo "make e2e   - run end-to-end (BDD) tests"

build:
	docker build -t $(IMAGE) .

run:
	docker run -d --rm --name $(IMAGE) -p $(PORT):$(PORT) $(IMAGE)

stop:
	-docker stop $(IMAGE)

test:
	go test -v -shuffle=on ./internal/endpoints/... ./advisor/...

e2e:
	go test -v ./internal/bdd/...
