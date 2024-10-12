FROM golang:1.23.2 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=build /app/api .

ENTRYPOINT [ "/app/api", "server" ]