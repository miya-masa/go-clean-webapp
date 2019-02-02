# go-clean-webapp

Clean Architecture GO

## Usage

1. `docker-compose up -d`
1. `go generate && make`
1. `./bin/webapp i`
1. `./bin/webapp s`

## API

```
http://localhost:8080
POST /accounts
{
  "first_name": "hoge",
  "last_name": "fuga"
}

http://localhost:8080
GET /accounts

http://localhost:8080
GET /accounts/{accountUUID}

http://localhost:8080
DELETE /accounts/{accountUUID}
```

