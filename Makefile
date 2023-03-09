build:
	cd cmd && go build -o server main.go

run:
	cd cmd && export HTTP_PORT=8080 && go run main.go

tidy:
	cd cmd && go mod tidy

test:
	cd cmd && go test -v ./...

compile:
	cd cmd && GOOS=linux GOARCH=amd64 go build -o server main.go

watch:
	cd cmd && ~/go/bin/reflex -s -r '\.go$$' make run