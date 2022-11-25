MONGO_SERVICE_NAME=mongo
GO_SERVICE_NAME=go

down:
	docker-compose down --volumes
up:
	docker-compose down 
	docker-compose up -d

exec_mongo:
	docker-compose exec $(MONGO_SERVICE_NAME) bash

exec_go:
	docker-compose exec $(GO_SERVICE_NAME) bash
