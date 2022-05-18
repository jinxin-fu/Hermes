FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata
    #&& apk add --no-cache dos2unix

WORKDIR /build

COPY . .
#RUN dos2unix run/run.sh
RUN go mod tidy
RUN go build -ldflags="-s -w" -o /app/hermes cmd/app/hermes.go
RUN #go build -ldflags="-s -w" -o /app/transform cmd/transform/transform.go


FROM ubuntu:18.04

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/hermes /app/hermes
COPY --from=builder /build/api/etc/hermes-api.yaml /app/hermes-api.yaml
#COPY --from=builder /app/transform /app/transform
#COPY --from=builder /build/rpc/transform/etc/transform.yaml /app/transform.yaml
#COPY --from=builder /build/run/run.sh /app/run.sh
#RUN chmod +x /app/run.sh
CMD ["./hermes","-f","hermes-api.yaml"]
#ENTRYPOINT ["/app/run.sh"]
