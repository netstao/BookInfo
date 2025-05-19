FROM golang:1.24.3-alpine3.21 AS builder
WORKDIR  /
COPY . .
RUN  [ ! -d "vendor" ] && go env -w GOPROXY=https://goproxy.cn,direct || echo "skipping GOPROXY setting, vendor directory exists"
RUN go build -o bk
FROM alpine:3.21-bookinfo-base
WORKDIR  /
COPY templates /templates
COPY static /static
COPY --from=builder /bk /
ENTRYPOINT [ "./bk", "serve" ]