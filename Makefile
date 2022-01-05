go:
	go run ./$(f) -env $(env)

docker.up:
	docker-compose -f $(f) up -d

docker.down:
	docker-compose -f $(f) down

docker.stop:
	docker-compose -f $(f) stop

docker.start:
	docker-compose -f $(f) start

docker.exec:
	docker exec -it $(it) bash