build:
	echo "building..."
	go build -o ./bin/news-api

run:
	echo "running..."
	./bin/news-api

clean:
	echo "cleaning..."
	rm -r ./bin

build_and_run: build run

docs:
	echo "generating docs..."
	swag init -g ./main.go