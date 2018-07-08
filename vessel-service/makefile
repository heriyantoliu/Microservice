build:
	docker build -t vessel-service .
run:
	docker run -p 50052:50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns \
	--name vessel-service \
	vessel-service
start:
	docker start -i vessel-service
stop:
	docker stop vessel-service
remove:
	docker rm vessel-service