FROM golang:latest

#ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/aynakeya/AynaAPI
COPY . $GOPATH/src/github.com/aynakeya/AynaAPI
RUN go build .

EXPOSE 8090
ENTRYPOINT ["./AynaAPI"]