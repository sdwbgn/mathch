FROM golang:alpine

WORKDIR /go/src/mathch
COPY . .
RUN apk add --update bash make gcc
RUN make
CMD ./bin/mathch $PORT
