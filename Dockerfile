FROM golang:alpine

RUN mkdir /files
COPY src/io/hello.go /files
WORKDIR /files

RUN go build -o /files/hello hello.go
ENTRYPOINT ["/files/hello"]
