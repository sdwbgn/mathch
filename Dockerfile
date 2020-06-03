FROM golang:1.14

WORKDIR /go/src/mathch
COPY . .
RUN make
EXPOSE 801
CMD ["./bin/mathch"]
