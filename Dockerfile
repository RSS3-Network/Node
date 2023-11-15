FROM golang:1.21.3-alpine AS builder

WORKDIR /root/rss3-node

COPY . .

RUN apk add --no-cache git make gcc libc-dev

RUN CGO_ENABLED=1 make build

FROM alpine:3.18.4 AS runner

WORKDIR /root/rss3-node

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /root/rss3-node/build/rss3-node .

ENTRYPOINT ["./rss3-node"]
