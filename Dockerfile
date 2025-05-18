FROM golang:1.22.4-slim
WORKDIR  /
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o bk
ENTRYPOINT [ "./bk", "server" ]