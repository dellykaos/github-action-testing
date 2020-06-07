export GO111MODULE        ?=on
export VAR_KUBE_NAMESPACE ?= latihan
export VERSION            ?= $(shell git show -q --format=%h)
export IMAGE              ?= dellynyman/latihan-go
export ENV                ?= development
export CONTEXT            ?= minikube

ODIR := deploy/_output

dep:
	go mod download

test:
	go test -cover -coverprofile=coverage.out ./...

coverage: test
	go tool cover -html=coverage.out

run:
	go run app/api/main.go

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/api/api app/api/main.go

build:
	docker build -t $(IMAGE):$(VERSION) -f ./deploy/api/Dockerfile . --build-arg VERSION=${VERSION}

push:
	docker push $(IMAGE):$(VERSION) && \
	docker rmi $(IMAGE):$(VERSION)

kubefile:
	$(foreach f, $(shell ls deploy/api/*.yml), \
		kubelize genfile --overwrite -c ./ -s api -e $(ENV) $(f) $(ODIR)/api/;)

apply:
	echo deploying "api" to environment "$(ENV)" && \
	kubectl apply -f $(ODIR)/api/service.yml --context=$(CONTEXT) && \
	kubectl apply -f $(ODIR)/api/deployment.yml --context=$(CONTEXT)

deploy: dep compile build push kubefile apply
