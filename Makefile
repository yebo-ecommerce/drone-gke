all:
	docker-compose run --rm web go build -o ./bin/drone-gke main.go
	docker-compose -f docker-compose.prod.yml build
	docker push yebodev/drone-gke
