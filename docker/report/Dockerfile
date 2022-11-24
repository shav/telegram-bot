FROM ubuntu:20.04

WORKDIR /app

COPY report ./

# Порт для сбора метрик
EXPOSE 8081

CMD ["/app/report", "-configPath", "/app/config.yaml"]