start-dev:
	docker-compose -f ./dev-script/docker-compose.db.yml up -d && docker-compose -f ./dev-script/docker-compose.yml up
	
stop-dev:
	@docker-compose -f ./dev-script/docker-compose.yml down --remove-orphan

start-prod:
	@docker-compose up --build

stop-prod:
	@docker-compose dowm