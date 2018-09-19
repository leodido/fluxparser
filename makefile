SHELL := /bin/bash

flux.go: flux.rl
	docker run --rm -v $(PWD):/tmp quay.io/influxdb/ragel:6.10 ragel -Z -o /tmp/$@ /tmp/$<
	#@gofmt -w -s $@
	#@sed -i '/^\/\/line/d' $@

.PHONY: build
build: flux.go

.PHONY:
clean: flux.go
	@rm -rf $^