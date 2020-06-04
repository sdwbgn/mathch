FROM golang:alpine

WORKDIR /go/src/mathch
COPY . .
RUN apk add --update bash make
RUN make
CMD ./bin/mathch $PORT
