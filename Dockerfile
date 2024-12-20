FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update

# build go app
RUN go mod download
RUN go build -o posts-log ./cmd/main.go

CMD ["./posts-log"]