FROM golang:1.9.2-alpine

RUN apk add --no-cache --update alpine-sdk

COPY . /go/src/github.com/PwC-ASD/dex-cli
RUN cd /go/src/github.com/PwC-ASD/dex-cli && make release-binary

FROM alpine

COPY --from=0 /go/bin/dex-cli /usr/local/bin/dex-cli

WORKDIR /

ENTRYPOINT ["dex-cli"]
CMD ["-h"]
