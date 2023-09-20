FROM golang:1.21.1-bullseye
WORKDIR /go/src/app

COPY ./. .

RUN go install -tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go mod download

CMD [ "go", "run", "main.go" ]
