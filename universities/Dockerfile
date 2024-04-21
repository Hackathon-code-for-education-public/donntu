FROM golang:1.22.0-alpine3.19 as builder

ENV GOPROXY=https://goproxy.io,direct
RUN apk update --no-cache
WORKDIR /app
COPY . /app
RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/service/app.go
RUN go build -mod=readonly -o migrator cmd/migrator/migrator.go

RUN ls -la && pwd && ls -la /app/db/migrations

FROM alpine

RUN apk update --no-cache
WORKDIR /app
COPY --from=builder /app/app ./app
COPY --from=builder /app/migrator ./migrator
COPY --from=builder /app/db/migrations ./migrations

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.12.1/wait /app/wait
RUN chmod +x /app/wait

RUN ls -la && pwd && ls -la /

CMD  ./wait \
    && ./migrator -dir /app/migrations \
    && ./app


