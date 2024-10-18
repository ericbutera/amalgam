FROM golang:1.23.2 as build

ARG APP_PATH="not-set"
ARG GO111MODULE="on"
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ${APP_PATH}/main.go

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=build /app/app .

ENTRYPOINT [ "/app/app" ]