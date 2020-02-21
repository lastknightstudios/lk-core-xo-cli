xo:
	@echo "---:xo: Creating xo symlink in /usr/local/bin"	
	@rm -f /usr/local/bin/xo
	
	@echo "---:xo: Building xo docker container"	
	@docker build . -t xo:latest
	ln -s $$PWD/cli/xo /usr/local/bin/xo

build: 
	@ls -la
	@pwd
	@go build -o xo

go:
	@export GOPATH=$(PWD)
	#@mkdir bin && mkdir src && mkdir pkg
	

path:
	@export PATH=$PATH:$(go env GOPATH)/bin

.PHONY: xo build