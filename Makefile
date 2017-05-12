.PHONY: build run frontend backend clean

build: frontend backend

run: build
	./server

frontend: static/bundles/main.js

static/bundles/main.js: $(wildcard frontend/*)
	cd frontend && npm run build && cd ..

backend: server

server: $(wildcard backend/*.go)
	cd backend && go build -o ../server

clean:
	-rm server
	-rm -r static/bundles
