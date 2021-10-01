FROM amazonlinux:2 as builder

ARG GO_FILE=go1.17.1.linux-amd64.tar.gz
ARG SRC_MODULE=github.com/mkaiho/gin-sample

ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

WORKDIR $GOPATH/src/$SRC_MODULE
COPY ./Makefile $GOPATH/src/$SRC_MODULE
COPY ./go.mod $GOPATH/src/$SRC_MODULE
COPY ./go.sum $GOPATH/src/$SRC_MODULE
COPY ./main.go $GOPATH/src/$SRC_MODULE

RUN yum install -y curl tar make gzip \
  && curl -OL https://golang.org/dl/${GO_FILE} \
  && tar -C /usr/local -xzf ${GO_FILE} \
  && rm ${GO_FILE} \
  && make build

FROM amazonlinux:2

RUN yum update -y && yum install -y shadow-utils \
  && useradd "ec2-user"

USER ec2-user
WORKDIR /ec2-user

COPY --from=builder /go/src/github.com/mkaiho/gin-sample/build ./build

CMD ["./build/main"]