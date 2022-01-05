go:
	go run ./$(f) -env $(env)

docker.up:
	docker-compose --env-file $(e) -f $(f) up -d

docker.down:
	docker-compose --env-file $(e) -f $(f) down

docker.stop:
	docker-compose --env-file $(e) -f $(f) stop

docker.start:
	docker-compose --env-file $(e) -f $(f) start

docker.exec:
	docker exec -it $(it) bash