build:
	docker build -t user-service .
run:
	docker run \
	--name user-service \
	-p 50053:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns \
	-e DB_NAME=postgres \
	-e DB_HOST=database \
	-e DB_PORT=5432 \
	-e DB_USER=postgres \
	-e DB_PASSWORD=postgres \
	user-service
start:
	docker start -i user-service
stop:
	docker stop user-service
remove:
	docker rm user-service