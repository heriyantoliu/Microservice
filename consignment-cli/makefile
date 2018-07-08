build:
	docker build -t consignment-cli .
run:
	docker run \
	-e MICRO_REGISTRY=mdns \
	--name consignment-cli \
	consignment-cli
start:
	docker start -i consignment-cli
stop:
	docker stop consignment-cli
remove:
	docker rm consignment-cli