REPO=jsdoitao/docker-ldap-passwd-webui
VER=1.1

.PHONY: all build push

all: build docker push

init:
	dep ensure
	
build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ldap-pass-webui main.go

docker:
	@echo "Building docker image"
	docker build -t ${REPO}:latest .

push: 
	@echo "Pushing to dockerhub"
	docker push ${REPO}:latest

clean:
	rm -rf ldap-pass-webui