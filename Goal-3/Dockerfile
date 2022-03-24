FROM golang:1.17.8-alpine3.15 as builder
ENV GO111MODULE on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV GOSUMDB=off
WORKDIR /go/cache
ADD go.mod .
RUN go mod download
WORKDIR  /go/src
COPY . .
RUN go get github.com/apache/pulsar-client-go/pulsar
RUN go get github.com/mkideal/cli
RUN go build -a -o runtime *.go

FROM alpine:3.15
WORKDIR /root
COPY . .
COPY --from=builder /go/src/runtime .
CMD ["./runtime"]