FROM alpine
WORKDIR /app
ADD bin bin
ENTRYPOINT bin/app