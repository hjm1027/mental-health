FROM golang:1.14.5
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.cn"
WORKDIR $GOPATH/src/github.com/mental-health/
COPY . $GOPATH/src/github.com/mental-health/
RUN make
EXPOSE 4096
CMD ["./main", "-c", "conf/config.yaml"]