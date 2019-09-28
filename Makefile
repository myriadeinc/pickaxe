up:
	docker-compose up

build:
	docker build -t myriade/pickaxe .

dev:
	docker build -f Dockerfile.dev -t myriade/pickaxe:dev .