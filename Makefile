.PHONY: up
up:
	docker-compose up --build --force-recreate -d

.PHONY: down
down:
	docker-compose down -v --rmi local