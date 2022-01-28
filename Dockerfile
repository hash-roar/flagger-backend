FROM golang:1.13 as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .


FROM alpine:latest

run mkdir /app

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]

