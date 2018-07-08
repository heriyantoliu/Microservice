run:
	docker run \
	-p \
	5432:5432 \
	-e MICRO_SERVER_ADDRESS=:5432 \
	-e MICRO_REGISTRY=mdns \
	--name postgres \
	postgres
stop:
	docker stop postgres
remove:
	docker rm postgres