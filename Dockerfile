FROM golang:1.16

ADD . /go/src/github.com/LexikonCMS/bucket-cdn

WORKDIR /go/src/github.com/LexikonCMS/bucket-cdn

RUN go install

ENTRYPOINT ["/go/bin/bucketCDN"]

EXPOSE 8080
