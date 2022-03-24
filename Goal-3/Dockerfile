FROM golang:1.17.8-alpine3.15 as builder
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.io/
ENV GOSUMDB=off
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR  /go/src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o runtime *.go

FROM alpine:3.15
ENV TOPIC=""
ENV SCRIPT=""
ENV INPUT=""
ENV OUTPUT=""
WORKDIR /root
COPY functions .
COPY --from=builder /go/src/runtime .
RUN chmod +x /root/runtime
ENTRYPOINT ["/root/runtime"]