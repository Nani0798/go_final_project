FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY . .

RUN go mod download

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o go_final_project ./cmd/scheduler/main.go

FROM alpine:3.18

ENV TODO_PORT=7540
ENV TODO_DBFILE=/app/storage/scheduler.db

WORKDIR /app

COPY --from=builder /app/go_final_project /app/
COPY --from=builder /app/.env .
COPY --from=builder /app/web ./web 

RUN mkdir -p /app/storage

CMD ["./go_final_project"]