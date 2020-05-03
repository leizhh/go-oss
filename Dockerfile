FROM golang:latest

RUN mkdir /storage 
RUN cd /storage && mkdir objects
RUN cd /storage && mkdir temp

COPY . /go/src/go-oss