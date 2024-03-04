up:
	echo "running..."
	docker-compose up -d postgres
	docker-compose up

test:
	echo "running tests..."
	docker-compose up -d postgres_test
	go test ./... -v

docs:
	echo "generating docs..."
	swag init -g ./main.go