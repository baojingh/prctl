VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
GIT_HASH = $(git show -s --format=%H)
GO_VERSION = $(go version)

all: clean build

push:
	git remote set-url origin  git@github.com:baojingh/prctl.git
	git pull
	git add .
	git commit -m "update"
	git push origin main

prehandle:clean
	sudo rm -rf /var/log/prctl/*
	sudo mkdir -p /var/log/prctl
	sudo chown -R ubuntu:ubuntu /var/log/prctl

run:prehandle
	go run main.go

install:build
	sudo cp build/prctl /usr/local/bin/prctl
	sudo chown -R ubuntu:ubuntu /usr/local/bin/prctl

build:prehandle
	mkdir build
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -ldflags="-s -w \
			-X 'github.com/baojingh/prctl/cmd.buildVersion=$(VERSION)' \
			-X 'github.com/baojingh/prctl/cmd.buildTime=$(CURRENT_TIME)'" \
			-v -o build/prctl

tidy:
	go mod tidy


clean:
	rm -rf build
	sudo rm -rf /usr/local/bin/prctl
