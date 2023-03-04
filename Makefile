BINARY=tf-view-app

default: build

build-server:
	go build -o ${BINARY} .

build-ui:
	npm install --prefix ui
	npm run build --prefix ui

build: build-ui build-server

run: build
	./${BINARY}

build-docker:
	docker build --target=prod . 

install: build
