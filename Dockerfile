FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -o fake-cli .

CMD ["./fake-cli"]

ENV CONFIG_PATH=/app/config.json
ENV PORT=8080
