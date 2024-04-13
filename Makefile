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
	sudo cp build/prctl /usr/local/bin/prctl
	sudo chown -R ${USER}:${USER} /usr/local/bin/prctl			

clean:
	rm -rf build
	sudo rm -rf /usr/local/bin/prctl

tidy:
	go mod tidy

deb-login:
	prctl login --url https://anc.com --username bob --password password

deb-logout:
	echo "TODO"

deb-down:
	sudo rm -rf /var/cache/apt/archives/*.deb
	sudo prctl  download  -i examples/deb.txt  -o /data/tmp/deb-pool

deb-ls:
	ls -lh /var/cache/apt/archives/

deb-upload:
	prctl  upload --architecture amd64 --component main --distribution bookworm \
				  --input /data/code/goproject/prctl/examples/deb-pool/

deb-del:
	echo "TODO"
