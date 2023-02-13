# Api del centro de informacion

### Docker config create db

```bash
docker run --name forum-go-app
    -e POSTGRES_PASSWORD={password}
    -e POSTGRES_USER={username}
    -e POSTGRES_DB={postgres}
    -p {puerto}
    -d postgres
```

### Execute docker postgresql

```bash
docker exec -it postgres-O9Cj psql -U postgres
```
