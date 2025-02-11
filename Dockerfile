FROM golang AS builder
WORKDIR /app
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build ./main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./server.go

FROM alpine
WORKDIR /server
COPY ./docker-entrypoint.sh ./
COPY --from=builder /app/main ./
COPY --from=builder /app/server ./
ENTRYPOINT ["/server/docker-entrypoint.sh"]
