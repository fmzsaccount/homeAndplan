#FROM golang:1.16 AS builder
#
#COPY . /src
#WORKDIR /src
#
#RUN GOPROXY=https://goproxy.cn make build
#
#FROM debian:stable-slim
#
#RUN apt-get update && apt-get install -y --no-install-recommends \
#		ca-certificates  \
#        netbase \
#        && rm -rf /var/lib/apt/lists/ \
#        && apt-get autoremove -y && apt-get autoclean -y
#
#COPY --from=builder /src/bin /app
#
#WORKDIR /app
#
#EXPOSE 8000
#EXPOSE 9000
#VOLUME /data/conf
#
#CMD ["./server", "-conf", "/data/conf"]

FROM test-quay.dc1-zh-cn.greeyun.com/official/alpine:latest

# We'll likely need to add SSL root certificates
RUN sed -i 's/dl-cdn.alpinelinux.org/mirror.gree.com/g' /etc/apk/repositories \
        && apk --no-cache add ca-certificates \
        && apk --no-cache add curl \
    	&& apk --no-cache add busybox-extras \
    	&& apk --no-cache add tzdata \
    	&& cp /usr/share/zoneinfo/Asia/Hong_Kong /etc/localtime \
    	&& echo "Asia/Hong_Kong" > /etc/timezone \
		&& mkdir -p /usr/local/app/configs \

WORKDIR /usr/local/app/cmd/gy_home/bin


COPY configs/config.yaml ./configs
COPY app .


EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./app", "-conf", "/data/conf"]
#CMD ["./app"]