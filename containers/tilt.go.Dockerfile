FROM alpine
WORKDIR /app
ADD bin bin
ENTRYPOINT /app/bin/app