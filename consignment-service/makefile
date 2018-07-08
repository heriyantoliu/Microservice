build:
	docker build -t consignment-service .
run:
	docker run -p 50051:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns \
	-e DB_HOST=datastore:27017 \
	--name consignment-service \
	consignment-service
start:
	docker start -i consignment-service
stop:
	docker stop consignment-service
remove:
	docker rm consignment-service