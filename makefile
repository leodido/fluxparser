SHELL := /bin/bash

flux.go: *.rl functions.go
	docker run -v $(PWD):/tmp -e CONTENT="`cat flux.rl`" --rm quay.io/influxdb/ragel:6.10 bash -c 'ragel -I /tmp -Z -o /dev/stdout <(echo "$$CONTENT")' > $@
	@gofmt -w -s $@
	@sed -i '/^\/\/line/d' $@

.PHONY: build
build: flux.go

.PHONY:
clean: flux.go
	@rm -rf $^