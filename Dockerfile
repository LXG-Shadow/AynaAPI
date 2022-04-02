FROM golang:alpine as build

RUN apk add build-base
#ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/github.com/aynakeya/AynaAPI
COPY . .

RUN go generate && go env && go build -buildvcs=false -o ayapi ./server/main

FROM alpine:latest

LABEL MAINTAINER="aynakeya@aynakeya.com"

WORKDIR /go/src/github.com/aynakeya/AynaAPI

COPY --from=build /go/src/github.com/aynakeya/AynaAPI ./

EXPOSE 8080
#ENV GIN_MODE=release
ENTRYPOINT ./ayapi -c=conf/conf_docker.ini