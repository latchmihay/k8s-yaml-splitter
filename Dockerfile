FROM golang:1.10.3-alpine3.8

RUN apk add --no-cache --update alpine-sdk bash

COPY . /go/src/github.com/mintel/k8s-yaml-splitter
WORKDIR /go/src/github.com/mintel/k8s-yaml-splitter
RUN make get && make 

FROM scratch
COPY --from=0 /go/src/github.com/mintel/k8s-yaml-splitter/bin/k8s-yaml-splitter /
ENTRYPOINT ["/k8s-yaml-splitter"]
CMD ["--help"]

