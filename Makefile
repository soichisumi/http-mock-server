go-build:
	go build -mod vendor -o exe .

docker-build:
	go mod vendor
	docker build -t soichisumi0/http-mock-server:$(TAG) .

docker-push:
	docker push soichisumi0/http-mock-server:$(TAG)