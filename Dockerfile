# Build stage only build binary file of golang
FROM golang:1.20-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
#RUN apk add curl
#RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080

# To define the default command to run when the container starts
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
