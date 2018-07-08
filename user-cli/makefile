build:
	docker build -t user-cli .
run:
	docker run \
	--name user-cli \
	-e MICRO_REGISTRY=mdns \
	user-cli
start:
	docker start -i user-cli
stop:
	docker stop user-cli
remove:
	docker rm user-cli