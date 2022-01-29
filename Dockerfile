FROM golang:1.13 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .


FROM  centos:latest


WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/config.yaml .

CMD ["/app/main"]

