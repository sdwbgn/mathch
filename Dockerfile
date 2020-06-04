FROM golang:alpine

WORKDIR /go/src/mathch
RUN apk add --update alpine-sdk
COPY . .
ENV GO_ENV production
RUN make
CMD ./bin/mathch $PORT
