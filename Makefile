VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
GIT_HASH = $(git show -s --format=%H)
GO_VERSION = $(go version)

all: clean install

push:
	git pull
	git add . || true
	git commit -m "update" || true
	git push origin main || true

install: clean
	mkdir build
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -ldflags="-s -w \
			-X 'github.com/baojingh/prctl/cmd.buildVersion=$(VERSION)' \
			-X 'github.com/baojingh/prctl/cmd.buildTime=$(CURRENT_TIME)'" \
			-v -o build/prctl		

clean:
	rm -rf build

tidy:
	go mod tidy

login:
	./build/prctl login --url https://jfrog.com/pool --repo repo  --username bob --password 123456
	ls -lh ~/.prctl
	cat ~/.prctl/config
	cat ~/.prctl/config | base64 -d
	echo ""

logout:
	./build/prctl logout
	ls -lh ~/.prctl

deb-down:
	sudo rm -rf examples/deb-pool
	./build/prctl deb download  -i examples/deb.txt  -o examples/deb-pool
	ls -lh examples/deb-pool

deb-upload:
	./build/prctl deb  upload --architecture amd64 --component main --distribution bionic \
				  --input examples/deb-pool/

deb-del:
	./build/prctl deb delete --all

pypi-down:
	sudo rm -rf examples/deb-pool
	./build/prctl pypi download  -i examples/deb.txt  -o examples/deb-pool
	ls -lh examples/deb-pool

pypi-upload:
	./build/prctl pypi  upload --architecture amd64 --component main --distribution bionic \
				  --input examples/deb-pool/

pypi-del:
	./build/prctl pypi delete --all
