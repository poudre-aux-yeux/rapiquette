FROM golang:latest as builder

WORKDIR /go/src/github.com/poudre-aux-yeux/rapiquette
COPY . .
RUN go get -u github.com/jteeuwen/go-bindata/...
RUN cd schema && go generate
RUN go get
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rapiquette

FROM alpine

WORKDIR /root/
COPY --from=builder /go/src/github.com/poudre-aux-yeux/rapiquette/rapiquette .
ENV GIN_MODE release
CMD ["./rapiquette"]

EXPOSE 3333
