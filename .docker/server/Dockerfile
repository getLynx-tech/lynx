FROM golang:1.26.1-alpine AS dev

WORKDIR /app

ENV PATH=/go/bin:$PATH

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ../../server/go.mod ../../server/go.sum ./
RUN go mod download

COPY ../../server .

ENTRYPOINT ["air", \
  "--build.cmd=go build -gcflags='all=-N -l' -o ./build/server ./cmd/main.go", \
  "--build.bin=./build/server", \
  "--build.delay=1000", \
  "--build.send_interrupt=true", \
  "--build.kill_delay=2000", \
  "--build.full_bin=dlv exec --accept-multiclient --headless --continue --log --listen=:40000 --api-version=2 ./build/server --" \
]

FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY ../../server/go.mod ../../server/go.sum ./
RUN go mod download

COPY ../../server .
ENV CGO_ENABLED=0
RUN go build -o server ./cmd/main.go

FROM golang:1.26.1-alpine AS runner

WORKDIR /app

COPY --from=builder /app /app
COPY ../../server/.env .env
COPY ../../server/internal/infrastructure/postgres/migrations ./internal/infrastructure/postgres/migrations
COPY --from=dev /go/bin/goose /usr/local/bin/goose

ENTRYPOINT [ "./server" ]