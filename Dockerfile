FROM golang:1.20-alpine AS builder
WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o brownout-controller

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/brownout-controller ./
CMD ["./brownout-controller"]