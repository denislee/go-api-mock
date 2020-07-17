FROM golang:1.14.4-alpine3.12

RUN apk upgrade && \
    apk update && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /var/cache/apk/*

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
