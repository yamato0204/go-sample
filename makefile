
.PHONY: up
up:
	docker compose up --build --watch

test:
	docker compose -f ./docker-compose.yml exec -T go sh -c "cd /app && go test -v ./tests/integration"


db-bash:
	docker compose exec -it db bash
