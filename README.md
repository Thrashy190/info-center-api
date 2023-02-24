# Api del centro de informacion

### Docker config create db

```bash
docker run --name info-center-api
    -e POSTGRES_PASSWORD=DB_PASSWORD
    -e POSTGRES_USER=DB_USER
    -e POSTGRES_DB=infocenter
    -p 55004:5432
    -d postgres
```

### Execute docker postgresql

```bash
docker exec -it info-center-api psql -U postgres
```

### Run Api

```bash
go run ./cmd/main.go
```
