FROM golang:1.12-alpine as builder

RUN apk add --no-cache alpine-sdk


ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

WORKDIR /go/src/github.com/alokic/sdkgen
COPY . .

RUN make build

###############################

FROM openapitools/openapi-generator-cli as sdkgen-image

RUN apk add --no-cache python3 && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --upgrade pip setuptools && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi && \
    if [[ ! -e /usr/bin/python ]]; then ln -sf /usr/bin/python3 /usr/bin/python; fi && \
    rm -r /root/.cache

RUN apk add --no-cache bash curl
RUN pip install yapf

ENV GOPATH /go
ENV PYTHON_POST_PROCESS_FILE "/usr/local/bin/yapf -i"
ENV ENV_TYPE "docker"

COPY --from=builder /go/bin/sdkgen /usr/bin/sdkgen
COPY --from=builder /go/bin/ctl /usr/bin/ctl
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

WORKDIR /go/src/github.com/alokic/sdkgen

ENTRYPOINT ["ctl"]