# ./Dockerfile-sender
FROM golang:1.20.2
RUN apt-get update -y && apt-get upgrade -y
RUN apt-get clean all
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.49.0
RUN mkdir /app
COPY ./app /app
WORKDIR /app

ENTRYPOINT go run /app/main.go
