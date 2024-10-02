FROM golang:1.22.5-alpine3.20 as builder

WORKDIR /build

COPY . .

FROM alpine:3.19.1

WORKDIR /dist

COPY --from=builder /build/server .

CMD ["./server"]