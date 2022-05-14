server:
	go run cmd/server/main.go

image:
	docker build -t api-gateway:latest .

containerup:
	docker run --name api-gateway --network neuromaps-network -p 8000:8000 api-gateway:latest

docker:
	docker start api-gateway

.PHONY: server containerup docker