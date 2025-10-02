DB_URL = postgres://postgres:1@localhost:5432/grownet_comments_db?sslmode=disable

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/comment.proto

migrate-up:
	migrate -path ./migrations/db -database ${DB_URL} up

migrate-down:
	migrate -path ./migrations/db -database ${DB_URL} down

migrate-force:
	migrate -path=migrations/db -database ${DB_URL} -verbose force 1