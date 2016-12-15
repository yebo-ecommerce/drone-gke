#
FROM golang:1.7

#
COPY . /go/src/github.com/yebo-ecommerce/drone-gke

#
WORKDIR /go/src/github.com/yebo-ecommerce/drone-gke

#
RUN go get -d -v

#
RUN go install -v
