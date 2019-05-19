#FROM alpine:latest
FROM debian:latest
RUN mkdir -p /app

WORKDIR /app

ADD user-cli /app/user-cli
ENTRYPOINT ["./user-cli"]
CMD ["./user-cli"]

