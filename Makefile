
docker-build:
	docker build --no-cache -t gcr.io/mchirico/mock-playground:test -f Dockerfile .

push:
	docker push gcr.io/mchirico/mock-playground:test

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/mchirico/mock-playground:test
