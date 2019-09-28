up:
	go run src/main.go

build:
	docker build -t myriade/pickaxe .

dev-build:
	docker build -t myriade/pickaxe:dev .