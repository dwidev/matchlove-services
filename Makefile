run: 
	@go run cmd/api/main.go
build: 
	@go build cmd/api/main.go

migrate:
	@go run cmd/api/main.go -migrate

seed:
	@go run cmd/api/main.go -seed

seedUser:
	@go run cmd/api/main.go -seed -user

seedMaster:
	@go run cmd/api/main.go -seed -master