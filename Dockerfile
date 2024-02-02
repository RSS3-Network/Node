FROM golang:1.21.4-alpine AS base

WORKDIR /root/node
RUN apk add --no-cache git make gcc libc-dev

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

COPY . .

FROM base AS builder

ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod/ \
    make build

FROM alpine:3.18.4 AS runner

WORKDIR /root/node

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /root/node/build/node .
COPY deploy/config.yaml /etc/rss3/node/config.yaml

ENTRYPOINT ["./node"]
