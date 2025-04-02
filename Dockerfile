FROM golang:1.23.2 AS builder

COPY robot-demo /src
WORKDIR /src

RUN GOPROXY=https://mirrors.aliyun.com/goproxy/,direct make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./robot-demo", "-conf", "/data/conf"]
