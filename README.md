# docker-networking

```shell
docker compose up --detach
docker compose run internal-only bash
curl example.com #works
curl github.com # 418
ssh git@github.com # DNS error
docker compose down -t 0
```
