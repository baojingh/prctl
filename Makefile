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

deb-login:
	./build/prctl login --url https://anc.com --repo myrepo --username bob --password password
	ls -lh ~/.prctl
	cat ~/.prctl/config
	cat ~/.prctl/config | base64 -d

deb-logout:
	./build/prctl logout
	ls -lh ~/.prctl

deb-down:
	sudo rm -rf examples/deb-pool
	go run main.go  download  -i examples/deb.txt  -o examples/deb-pool
	ls -lh examples/deb-pool

	# ./build/prctl  download  -i examples/deb.txt  -o examples/deb-pool

deb-ls:
	ls -lh /var/cache/apt/archives/

deb-upload:
	./build/prctl  upload --architecture amd64 --component main --distribution bookworm \
				  --input /data/code/goproject/prctl/examples/deb-pool/

deb-del:
	echo "TODO"
