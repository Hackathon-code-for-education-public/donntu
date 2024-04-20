BIN_DIR=bin

build:
	mkdir -p $(BIN_DIR) #
	go mod tidy
	go build -o $(BIN_DIR)/ -v ./cmd/service

gen:
	protoc  \
	 	--go_out=. --go-grpc_out=. \
		-I ../proto ../proto/students.proto
	wire ./internal/app

deploy:
	docker compose build
	docker compose down
	docker compose up -d

clean:
	rm -r $(BIN_DIR) api #