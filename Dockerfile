# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.16-alpine

ENV GOBIN /usr/local/bin
WORKDIR /go/src/app
RUN go mod init
RUN go get -u github.com/go-http-utils/logger

COPY main.go serve.go

#RUN go get -d -v ./main.go
RUN go install -v ./serve.go

# Run the outyet command by default when the container starts.
CMD /usr/local/bin/serve

# Document that the service listens on port 8080.
EXPOSE 80

