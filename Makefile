SERVICE_NAME= mysql

down:
	docker-compose down --volumes
	
up:
	docker-compose down --volumes
	docker-compose up \
	--build \
	-d

exec:
	docker-compose exec $(SERVICE_NAME) bash
