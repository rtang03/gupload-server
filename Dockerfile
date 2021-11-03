FROM golang:1.15.1-alpine3.12 AS builder

LABEL stage=builder

RUN apk add --no-cache gcc libc-dev

WORKDIR /workspace

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -i -o build/gupload main.go

FROM nginx:1.21.3-alpine AS final
WORKDIR /var/gupload
VOLUME /var/gupload/cert

COPY --from=builder /workspace/build/gupload .
COPY --from=builder /workspace/README.md .
COPY --from=builder /workspace/cert ./cert
RUN touch /usr/share/nginx/html/.wellknown \
    && ln -s /var/gupload/gupload /usr/share/nginx/html/gupload \
    && ls /usr/share/nginx/html

CMD ["sh", "-c", "./gupload serve --key ./cert/tls.key --certificate ./cert/tls.crt"]

