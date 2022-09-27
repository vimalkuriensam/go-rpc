#  Binary file name
BROKER_BINARY=brokerApp

#  Starting all docker containers from docker compose file
up:
	@echo "Starting all docker images..."
	docker-compose -f ./project/docker-compose.yml up -d
	@echo "Done"

# Stopping all running containers and rebuilding it
up_build:
	@echo "Stopping all running containers"
	docker-compose -f ./project/docker-compose.yml down
	@echo "All containers stopped"
	@echo "Building all images ..."
	docker-compose -f ./project/docker-compose.yml up --build -d
	@echo "Done"

# Stopping all running containers
down:
	@echo "Stopping all running containers"
	docker-compose -f ./project/docker-compose.yml down
	@echo "Done"

# Creating the binary file for the item-service
build_item_search:
	@echo "Building item search binary"
	cd ./item-service && go build -o ../bin/${BROKER_BINARY} ./cmd/api/.
	@echo "Done"