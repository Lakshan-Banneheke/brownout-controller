FROM golang:1.20-alpine AS builder
WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN env GOOS=linux GOARCH=arm GOARM=7 go build

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/brownout-controller ./
CMD ["./brownout-controller"]