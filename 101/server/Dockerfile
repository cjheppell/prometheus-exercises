FROM golang:1.15-alpine

WORKDIR /go/src/app
ADD go.mod go.mod

RUN go mod download

ADD . .

RUN go install -v ./...

CMD ["server"]