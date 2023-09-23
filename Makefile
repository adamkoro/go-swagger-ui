build:
	cd swagger-ui && go build -o server main.go

run:
	cd swagger-ui && export HTTP_PORT=8080 && go run main.go

tidy:
	cd swagger-ui && go mod tidy

test:
	cd swagger-ui && go test -v ./...

compile:
	cd swagger-ui && GOOS=linux GOARCH=amd64 go build -o server main.go

watch:
	cd swagger-ui && ~/go/bin/reflex -s -r '\.go$$' make run