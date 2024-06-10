FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o build cmd/networkator/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build /build

EXPOSE 8080
ENTRYPOINT ["./build"]