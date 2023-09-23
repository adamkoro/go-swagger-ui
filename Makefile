build:
	cd swagger-ui && go build -o swagger-ui swagger-ui.go

run:
	cd swagger-ui && export HTTP_PORT=8080 && go run swagger-ui.go

tidy:
	cd swagger-ui && go mod tidy

test:
	cd swagger-ui && go test -v ./...
