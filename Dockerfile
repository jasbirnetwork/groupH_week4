FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/myapp .

CMD [ "./myapp" ]