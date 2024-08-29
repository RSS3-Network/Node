FROM ghcr.io/rss3-network/go-image/go-builder:main-4782eed AS base

WORKDIR /root/node

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN wget -O etherface.tar.gz https://storage.googleapis.com/rss3-etherface/etherface.tar.gz && \
    tar -xzf etherface.tar.gz && \
    rm etherface.tar.gz

COPY . .

FROM base AS builder

ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod/ \
    make build

FROM ghcr.io/rss3-network/go-image/go-runtime:main-4782eed AS runner

WORKDIR /root/node

COPY --from=builder /root/node/etherface /root/node/etherface

COPY --from=builder /root/node/build/node .
COPY deploy/default/config.default.yaml /etc/rss3/node/config.yaml

ENTRYPOINT ["./node"]
