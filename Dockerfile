FROM golang:alpine

WORKDIR /go/src/mathch
COPY . .
RUN apk add --update alpine-sdk
RUN make
CMD ./bin/mathch $PORT
