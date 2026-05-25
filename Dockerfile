FROM golang:1.25-alpine AS builder

LABEL maintainer="Codee"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/release

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add tzdata

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.10 && swag init  --parseDependency=true
RUN CGO_ENABLED=0 GOOS=linux go build -p 1 -ldflags="-w -s" -a -installsuffix cgo -o golang_per_day .

FROM alpine:latest

COPY --from=builder /go/release/golang_per_day /

COPY --from=builder /go/release/configs /configs

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8080

CMD ["/golang_per_day","server","-c", "/configs/config.yaml"]
