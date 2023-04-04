FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /usr/src/app

COPY go.mod .

RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/app ./...

FROM alpine:3.17.3

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD [ "app" ]
