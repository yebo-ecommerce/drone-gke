all:
	docker-compose run --rm web go build
	docker-compose -f docker-compose.prod.yml build
	docker push yurifl/drone-gke
