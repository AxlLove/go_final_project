FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o scheduler .

FROM ubuntu:latest

WORKDIR /app
COPY --from=builder /app/scheduler .
COPY web ./web

ENV TODO_PORT=7540
ENV TODO_DBFILE=/data/scheduler.db

EXPOSE 7540

CMD ["./scheduler"]
