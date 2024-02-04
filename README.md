# Api del centro de informacion

[App web repo](https://github.com/Thrashy190/info-center-app)

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

### Run main

```bash
go run .
```
