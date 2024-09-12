
.PHONY: up
up:
	docker compose up --build --watch

db-bash:
	docker compose exec -it db bash
