# Compile
FROM golang:1.18.10-alpine3.16 AS builder

RUN mkdir /app
ADD . /app/
WORKDIR /app

ENV GOPROXY https://goproxy.cn,direct

RUN go build -o tokenizer .


# run
FROM alpine:3.16

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

RUN mkdir /app \
    && apk upgrade \
    && apk add --no-cache bash tzdata \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone

WORKDIR /app
COPY --from=builder /app/tokenizer .

CMD ["./tokenizer"]