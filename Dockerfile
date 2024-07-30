FROM golang:1.22-alpine AS builder

WORKDIR /usr/src/app
RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ./app/ ./

RUN go mod download && go mod verify
RUN go build -o /usr/local/bin/app ./main.go

FROM alpine

COPY --from=builder /usr/src/app ./
