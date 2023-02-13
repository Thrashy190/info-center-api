# Api del centro de informacion

### Docker config create db

```bash
docker run --name forum-go-app
    -e POSTGRES_PASSWORD=posxstgrespw
    -e POSTGRES_USER=postgres
    -e POSTGRES_DB=postgres
    -p 55002:5432
    -d postgres
```

### Execute docker postgresql

```bash
docker exec -it postgres-O9Cj psql -U postgres
```
