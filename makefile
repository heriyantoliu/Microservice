build:
	docker-compose build
up:
	docker-compose up
remove:
	docker rm vessel-service
	docker rm consignment-service
	docker rm consignment-cli
	docker rm datastore
	docker rm database
	docker rm user-cli
	docker rm user-service
clean:
	docker image prune