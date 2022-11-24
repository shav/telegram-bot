FROM ubuntu:20.04

WORKDIR /app

COPY bot ./

COPY ca-certificates.crt /etc/ssl/certs/

# Порт для сбора метрик
EXPOSE 8080

CMD ["/app/bot", "-configPath", "/app/config.yaml"]