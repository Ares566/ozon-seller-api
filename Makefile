dev: 
	go run cmd/main.go
docker:
	docker build . -t ozoneapi
	docker run -it --rm ozoneapi
