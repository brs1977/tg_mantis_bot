FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

WORKDIR /build

ADD go.mod .

ADD go.sum .

COPY . .

RUN go build

FROM alpine

WORKDIR /build

COPY --from=builder /build/mantis /build/mantis

CMD ["./mantis"]