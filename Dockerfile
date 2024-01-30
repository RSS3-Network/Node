FROM golang:1.21.4-alpine AS builder

WORKDIR /root/node

COPY . .
RUN apk add --no-cache git make gcc libc-dev

RUN CGO_ENABLED=1 make build

FROM alpine:3.18.4 AS runner

WORKDIR /root/node

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /root/node/build/node .

ENTRYPOINT ["./node"]
