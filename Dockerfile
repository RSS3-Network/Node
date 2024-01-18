FROM golang:1.21.3-alpine AS builder

WORKDIR /root/rss3-node

COPY . .

ARG GH_USER
ARG GH_TOKEN

ENV GOPRIVATE="github.com/naturalselectionlabs/rss3-global-indexer"
ENV GH_USER=$GH_USER
ENV GH_TOKEN=$GH_TOKEN
RUN git config --global url."https://${GH_USER}:${GH_TOKEN}@github.com".insteadOf "https://github.com"

RUN apk add --no-cache git make gcc libc-dev

RUN CGO_ENABLED=1 make build

FROM alpine:3.18.4 AS runner

WORKDIR /root/rss3-node

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /root/rss3-node/build/rss3-node .

ENTRYPOINT ["./rss3-node"]
