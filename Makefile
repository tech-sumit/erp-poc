.PHONY: clean start_stack stack_logs

clean:
	rm -rf ./build

build: clean
	export GO111MODULE=on
	env GOOS=darwin go build -o main.go -o build/gateway

build_docker: clean
#	export GO111MODULE=on
#	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main.go -o build/gateway
	docker build -t gcr.io/formal-triode-302008/erp-gateway .

docker_push: build_docker
	docker push gcr.io/formal-triode-302008/erp-gateway

start_docker: build_docker
	docker run -p 0.0.0.0:8080:8080 gcr.io/formal-triode-302008/erp-gateway

start_stack: clean
	export GO111MODULE=on
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main.go -o build/gateway
	docker-compose up -d --build

stop_stack:
	docker-compose down

stack_logs:
	docker-compose logs -f

start_local: clean build
	build/gateway
