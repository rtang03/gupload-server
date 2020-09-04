FROM golang:1.15.1-alpine3.12 AS builder

LABEL stage=builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /workspace

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -i -o build/gupload main.go

FROM alpine:3.12 AS final

WORKDIR /var/gupload

VOLUME /var/gupload/uploaded

COPY --from=builder /workspace/build/gupload .
COPY --from=builder /workspace/README.md .
COPY --from=builder /workspace/cert .

CMD ["sh", "-c", "sleep 48h"]

