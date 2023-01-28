FROM golang:alpine AS builder

RUN apk update  \
 && apk add --no-cache git

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY . .

RUN go get -v

RUN go test \
 && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /listener

FROM scratch

COPY --from=builder /etc/group  /etc/passwd /etc/
COPY --from=builder /listener /go/bin/listener

USER appuser:appuser

ENTRYPOINT ["/go/bin/listener"]
