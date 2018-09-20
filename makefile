SHELL := /bin/bash

flux.go: flux.rl
	docker run -e CONTENT="`cat $<`" --rm quay.io/influxdb/ragel:6.10 bash -c 'ragel -Z -G2 -e -o /dev/stdout <(echo "$$CONTENT")' > $@
	@gofmt -w -s $@
	@sed -i '/^\/\/line/d' $@

.PHONY: build
build: flux.go

.PHONY:
clean: flux.go
	@rm -rf $^