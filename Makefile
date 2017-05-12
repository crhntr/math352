.PHONY: build run frontend backend clean

build: frontend backend

run: build jwt_token_key.priv
	./server

frontend: static/bundles/main.js

static/bundles/main.js: $(wildcard frontend/*)
	cd frontend && npm run build && cd ..

backend: server

server: $(wildcard backend/*.go)
	cd backend && go build -o ../server

jwt_token_key.priv:
	openssl genrsa -out backend/jwt_token_key.priv 2048
	openssl rsa -in backend/jwt_token_key.priv -pubout > backend/jwt_token_key.pub

clean:
	-rm server
	-rm -r static/bundles
