FROM registry.suse.com/bci/golang:1.21 as builder
WORKDIR /go/src/builder
COPY swagger-ui/ .
RUN go build -ldflags="-s -w" -o swagger-ui

FROM registry.suse.com/bci/bci-micro:15.5
USER root
RUN echo "user:x:10000:10000:user:/home/user:/bin/bash" >> /etc/passwd && mkdir -p /home/user/static
USER user
WORKDIR /home/user
COPY --from=builder /build/swagger-ui/swagger-ui ./swagger-ui
ENV GIN_MODE=release \
    HTTP_PORT=8080 \
    STATIC_FILE_PATH=/home/user/static \
    SWAGGER_URL=https://raw.githubusercontent.com/neuvector/neuvector/main/controller/api/apis.yaml
EXPOSE ${HTTP_PORT}
ENTRYPOINT ["./swagger-ui"]