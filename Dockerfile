FROM golang:1.17-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN apk update && apk upgrade && apk add bash
COPY --from=builder /build/dtn-server /app/
COPY wait-for-it.sh /app/
WORKDIR /app
CMD ["./dtn-server"]