FROM golang:1.12-alpine as builder

RUN apk add --no-cache alpine-sdk


ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

WORKDIR /go/src/github.com/alokic/sdkgen
COPY . .

RUN make build

###############################

FROM alpine:edge as sdkgen-image

RUN apk add --no-cache bash curl

ENV GOPATH /go

COPY --from=builder /go/bin/sdkgen /usr/bin/sdkgen
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

WORKDIR /go/src/github.com/alokic/sdkgen

CMD [ "sdkgen" ]