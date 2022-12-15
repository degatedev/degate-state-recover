FROM golang:1.20rc1

ENV LC_ALL C.UTF-8
ENV LANG C.UTF-8

RUN apt-get update && apt-get install -y git gcc vim
RUN cd /go/src/ && git clone https://github.com/degatedev/degate-state-recover.git
RUN cd /go/src/degate-state-recover && go mod download