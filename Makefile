VERSION = 0.0.1
CURRENT_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
GIT_HASH = $(git show -s --format=%H)
GO_VERSION = $(go version)

all: clean install

push-hub:
	git  remote set-url origin git@github.com:baojingh/prctl.git
	git pull
	git add . || true
	git commit -m "update" || true
	git push origin main || true


push:
	git  remote set-url origin git@code.siemens.com:baojing.he/prctl.git
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
	# ./build/prctl login --url https://jfrog.com/pool --repo repo  --username bob --password 123456
	./build/prctl login --url https://ff.rtf-alm.ff.cloud --repo ff-dev-go-ff  --username ee --password fd
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

deb-list:
	./build/prctl deb list --all

pypi-down:
	sudo rm -rf examples/pypi-pool
	./build/prctl pypi download  -i examples/requirements.txt  -o examples/pypi-pool
	ls -lh examples/pypi-pool

pypi-upload:
	./build/prctl pypi  upload --input examples/pypi-pool/

pypi-del:
	./build/prctl pypi delete --all

pypi-list:
	./build/prctl pypi list --all


go-down:
	go clean -modcache
	# ls -lh /home/${USER}/go/pkg
	# ll /root/go/pkg/mod
	# sudo rm -rf examples/go-pool
	./build/prctl go download  -i examples/  -o examples/go-pool
	# ls -lh /home/${USER}/go/pkg

go-upload:
	./build/prctl go  upload --input /root/go/pkg/mod/cache/

go-del:
	./build/prctl go delete --all

go-list:
	./build/prctl go list --all


