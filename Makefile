
.PHONY: nats

nats:
	go build -o receiver ./nats/main.go
	./receiver

.PHONY: server

server:
	go build -o server ./cmd/main.go
	./server
 
 
.PHONY: publisher

publisher:
	go build -o publisher ./cmd/publisher/publisher.go
	./publisher

.PHONY: up
up:
	docker-compose up -d


.DEFAULT_GOAL := build