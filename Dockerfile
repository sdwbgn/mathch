FROM golang:alpine

WORKDIR /go/src/mathch
COPY . .
RUN make
EXPOSE 801
CMD ./bin/mathch $PORT
