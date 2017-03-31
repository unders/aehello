VERSION=0.0.2
PROG=helloworld
IMAGE=$(PROG):$(VERSION)
MAIN_VERSION="main.Version=v$(VERSION)"
BUILDSTAMP="main.Buildstamp=$(shell date -u '+%Y-%m-%dT%I:%M%p')"
GITHASH="main.Githash=$(shell git rev-parse HEAD)"
LDFLAGS=-ldflags "-X $(MAIN_VERSION) -X $(BUILDSTAMP) -X $(GITHASH)"

http?=":10080"
httpAddr=$(http)
httpPort=$(shell echo "$(http)" | sed 's/\://g')

# appengine projectID
# projectID="your project id"
SERVICE_NAME=default
REMOTE_IMAGE=gcr.io/$(projectID)/$(IMAGE)

.PHONY: help
help:
	@echo ""
	@echo "** $(PROG) v$(VERSION)   [ project: $(projectID) ] **"
	@echo ""
	@echo "Commands:"
	@echo "    make build         # builds a docker image: $(IMAGE)"
	@echo "    make ae/push       # push $(IMAGE) to appeninge"
	@echo "    make ae/deploy     # deploy and start $(REMOTE_IMAGE) as version: dev"
	@echo "    make ae/start      # start $(REMOTE_IMAGE) as version: dev"
	@echo "    make ae/stop       # stop $(REMOTE_IMAGE) as version: dev"
	@echo "    make ae/list       # list services on appengine"
	@echo "    make ae/tail       # tail log on appengine"
	@echo "    make test          # runs the tests"
	@echo "    make lint          # runs a linter on the source code"
	@echo "    make list          # shows containers and images"
	@echo "    make start         # starts container"
	@echo "    make stop          # stops container"
	@echo "    make run           # runs the go binary"
	@echo "    make log           # show log for running container"
	@echo "    make tail          # tails log for running container"
	@echo "    make check         # do requests against the running container"
	@echo "    make clean         # stops the container and removes it and its image"
	@echo ""
	@echo "Args (start, run, check):"
	@echo "    http               # http=:10180"
	@echo ""
	@echo "Examples:"
	@echo "    make start"
	@echo "    http=:10180 make start"
	@echo "    http=:10180 make check"
	@echo "    http=:10180 make run"
	@echo "    make push"
	@echo ""

.PHONY: build
build: test lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS)
	@file $(PROG)
	docker build -t $(IMAGE) .
	@rm $(PROG)
	@docker images $(IMAGE)

.PHONY: ae/push
ae/push: #build
	docker tag $(IMAGE) $(REMOTE_IMAGE)
	gcloud docker -- push $(REMOTE_IMAGE)

.PHONY: ae/deploy
ae/deploy:
	gcloud app deploy app-dev.yaml --image-url $(REMOTE_IMAGE) --version dev

.PHONY: ae/start
ae/start:
	gcloud app versions start dev --service=$(SERVICE_NAME)

.PHONY: ae/stop
ae/stop:
	gcloud app versions stop dev --service=$(SERVICE_NAME)
.PHONY: ae/stop

.PHONY: ae/list
ae/list:
	gcloud app versions list --service=$(SERVICE_NAME)

.PHONY: ae/tail
ae/tail:
	gcloud app logs tail --service=$(SERVICE_NAME)

.PHONY: test
test:
	@go test

.PHONY: lint
lint: install
	gometalinter \
	--enable-all \
	--deadline=100s \
	--line-length=100 \
	./...

.PHONY: list
list:
	docker ps | grep $(PROG) || true
	docker ps -a | grep $(PROG) || true
	docker images $(PROG)

.PHONY: start
start:
	docker run -p $(httpPort):8080 --name $(PROG) -d $(IMAGE)
	docker ps -l | grep $(IMAGE) | awk '{print $$1}' | xargs docker logs

.PHONY: stop
stop:
	docker ps | grep $(PROG) | awk '{print $$1}' | xargs docker stop
	docker ps -al | grep $(IMAGE) | awk '{print $$1}' | xargs docker logs --tail 20
	@docker ps -a  | grep $(IMAGE) | awk '{print $$1}' | xargs docker rm

.PHONY: install
install:
	@go install

.PHONY: run
run: install
	$(PROG) -http=$(httpAddr) -l

.PHONY: log
log:
	docker ps -al | grep $(IMAGE) | awk '{print $$1}' | xargs docker logs

.PHONY: tail
tail:
	docker ps -al | grep $(IMAGE) | awk '{print $$1}' | xargs docker logs -f --tail 10

.PHONY: check
check:
	curl http://localhost:$(httpPort)
	curl http://localhost:$(httpPort)/what?

.PHONY: clean
clean:
	docker ps -a | grep $(PROG) | awk '{print $$1}' | xargs docker stop
	docker ps -a | grep $(PROG) | awk '{print $$1}' | xargs docker rm
	docker images $(PROG) -q | xargs docker rmi -f
	docker images | grep none | awk '{print $$3}' | xargs docker rmi