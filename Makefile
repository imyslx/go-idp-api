#! /usr/bin/make

REPO:=github.com/imyslx/go-idp-api

install:
	@go get github.com/couchbase/gocb
	@go get github.com/rs/zerolog/log
	@go get gopkg.in/yaml.v2

gen: clean generate

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client  -d $(REPO)/design