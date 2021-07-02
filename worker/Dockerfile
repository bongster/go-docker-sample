FROM golang:1.16.4-alpine

RUN apk update -qq

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /bin/app ./src

EXPOSE 8080

CMD ["/bin/app"]
