
up:
	docker-compose up

dev:
	docker build -t myriade/pickaxe:dev -f Dockerfile.dev .

build:
	docker build -t myriade/pickaxe .

