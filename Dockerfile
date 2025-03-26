FROM golang:1.24-alpine3.21

RUN addgroup -g 10001 logtester && \
    adduser -u  10001 -G logtester -D logtester

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o log-tester main.go

EXPOSE 8080
CMD ["./log-tester"]