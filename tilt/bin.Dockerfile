# cannot use scratch as it is read only
FROM alpine
WORKDIR /app
ADD bin bin
ENTRYPOINT /app/bin/app