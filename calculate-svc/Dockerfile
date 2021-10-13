FROM golang:1.15-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

WORKDIR /app

COPY src/go.mod ./

RUN go mod download

COPY src/ .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
