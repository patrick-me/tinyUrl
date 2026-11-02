# tinyUrl
A simple tinyUrl

---

## Simple usage

#### Create URL

> curl -d '{"url":"https://ya.ru"}' localhost/api/v1/shortUrl && echo
```
URL created!
{"shortUrl":"http://localhost/Ch9Lgr9kzW"}

```

#### Get URL

> curl -L http://localhost/Ch9Lgr9kzW
```
... Redirect to https://ya.ru
```

#### Default expiration
- 720 hours
---

## Docker usage

#### Compose
> docker-compose up

To rebuild use this cmd
> docker-compose -f docker-compose.yml up -d --build

#### Logs
Each container keeps at most three 10 MB log files. The limit is set when a
container is created, so existing containers pick it up after
`docker compose up -d` recreates them.

#### Upgrading from the Redis 8 image
Redis now runs on Valkey 8 (`valkey/valkey:8-alpine`) with a new volume,
`valkey-data`. Earlier versions built `redis:alpine` (Redis 8) and kept data in
`redis-data`. Valkey cannot read Redis 8 data files, so copy the links over
once:

1. Start the new version. Redis comes up empty, and old links return 404
   until step 3 is done.
   ```
   docker compose up -d
   ```
2. Start the old data in a temporary container. The volume is named
   `<project>_redis-data` (`docker volume ls`); the project name defaults to
   the folder name.
   ```
   docker run -d --name tinyurl_redis_old -v tinyurl_redis-data:/data redis:8-alpine redis-server --appendonly yes
   ```
3. Copy the keys. The script ends with `ok`, or exits with an error if any
   key did not make it.
   ```
   scripts/copy-redis-keys.sh tinyurl_redis_old tinyurl_redis
   ```
4. Remove the temporary container. Once the links work, remove the old volume
   and images too.
   ```
   docker rm -f tinyurl_redis_old
   docker volume rm tinyurl_redis-data
   docker image rm tinyurl-redis redis:8-alpine
   ```
