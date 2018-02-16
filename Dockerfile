FROM golang:latest as builder

WORKDIR /go/src/github.com/poudre-aux-yeux/rapiquette
COPY . .
RUN go get -u github.com/magefile/mage
RUN mage getgogenerate
RUN mage schema
RUN mage getdep
RUN mage vendorci
RUN mage buildci

FROM alpine

WORKDIR /root/
COPY --from=builder /go/src/github.com/poudre-aux-yeux/rapiquette/rapiquette .
ENV GIN_MODE release
CMD ["./rapiquette"]

EXPOSE 3333
