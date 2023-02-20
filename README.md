# Api del centro de informacion

### Docker config create db

```bash
docker run --name info-center-api
    -e POSTGRES_PASSWORD=postgrespw
    -e POSTGRES_USER=postgres
    -e POSTGRES_DB=postgres
    -p 55004:5432
    -d postgres
```

### Execute docker postgresql

```bash
docker exec -it info-center-api psql -U postgres
```
