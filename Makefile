export WEB_ENV := test

install:
	go mod download
	go mod tidy
	WEB_ENV=dev go run ./cmd/main.go

app_test:
	go mod download
	go mod tidy
	WEB_ENV=test go test ./test/config_test.go
	WEB_ENV=test go test ./test/wirefire_test.go