FROM rust:1.68-alpine3.17 as builder

WORKDIR /usr/src/app

COPY . .

RUN apk add --no-cach musl-dev && cargo install --path .

FROM alpine:3.17.3

COPY --from=builder /usr/local/cargo/bin/app /usr/local/bin/app

CMD [ "app" ]
