## Архитектура

Требования и архитектура приложения:
https://miro.com/app/board/uXjVPKgXgo4=/

Количественные характеристики:
https://disk.yandex.ru/i/3bIyUDsYWr8Nyw

## Tracing

`make tracing`

Jaeger: http://127.0.0.1:16686/

## Metrics

`make metrics`

Prometheus: http://127.0.0.1:9090/

Grafana: http://127.0.0.1:3000/ (admin/admin)

При первом логине в Графану она попросит установить новый пароль, ставим.

Заходим в шестеренку слева, выбираем Data sources, добавляем Prometheus, адрес `http://prometheus:9090`

## Caching

`make caching`

Redis is available on redis://127.0.0.1:6379