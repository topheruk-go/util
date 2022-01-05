go:
	go run ./$(f) -env $(env)

docker.up:
	docker-compose --env-file $(d).env -f $(d)docker-compose.yaml up -d

docker.down:
	docker-compose --env-file $(d).env -f $(d)docker-compose.yaml down

docker.stop:
	docker-compose --env-file $(d).env -f $(d)docker-compose.yaml stop

docker.start:
	docker-compose --env-file $(d).env -f $(d)docker-compose.yaml start

docker.exec:
	docker exec -it $(it) bash