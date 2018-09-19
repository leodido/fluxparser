SHELL := /bin/bash

flux.go: flux.rl
	ragel -Z -G2 -e -o $@ $<
	@gofmt -w -s $@
	@sed -i '/^\/\/line/d' $@

.PHONY: build
build: flux.go

.PHONY:
clean: flux.go
	@rm -rf $^