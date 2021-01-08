FROM golang:1.15

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/src/github.com/Le0tk0k/go-flat-api-sample

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 8080
CMD ["go", "run", "."]