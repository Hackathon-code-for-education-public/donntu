migrate.new:
	migrate create -ext sql -dir ./db/migrations -seq $(name)
gen:
	protoc -I ../proto ../proto/*.proto --go_out=. --go-grpc_out=.
	wire ./internal/app
