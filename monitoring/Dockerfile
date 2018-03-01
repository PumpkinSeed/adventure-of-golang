FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/github.com/PumpkinSeed/adventure-of-golang/monitoring
WORKDIR /code/src/github.com/PumpkinSeed/adventure-of-golang/monitoring
RUN go build -o monitoring

ENTRYPOINT /code/src/github.com/PumpkinSeed/adventure-of-golang/monitoring/monitoring