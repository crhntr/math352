.PHONY: build run frontend backend clean

build: frontend backend

run: build
	./math352

frontend: static/bundles/main.js

static/bundles/main.js: $(wildcard frontend/*)
	npm run build

backend: math352

math352: $(wildcard *.go)
	go build

clean:
	-rm math352
	-rm -r static/bundles
