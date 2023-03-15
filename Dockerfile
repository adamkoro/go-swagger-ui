FROM harbor.adamkoro.com/bci/bci-micro:15.4
WORKDIR /home/user
RUN mkdir static
# /build mount comes from CI temp directory
ADD /build/swagger-ui ${WORKDIR}
ENV GIN_MODE=release \
    HTTP_PORT=8080 \
    STATIC_FILE_PATH=/home/user/static \
    SWAGGER_URL=https://raw.githubusercontent.com/neuvector/neuvector/main/controller/api/apis.yaml
RUN echo "user:x:10000:10000:user:/home/user:/bin/bash" >> /etc/passwd && chown -R user /home/user/ && chmod u+x swagger-ui
EXPOSE ${HTTP_PORT}
USER user
ENTRYPOINT ["./swagger-ui"]