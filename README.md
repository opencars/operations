# operations

[![Go Report Card](https://goreportcard.com/badge/github.com/opencars/operations)](https://goreportcard.com/report/github.com/opencars/operations)

## Development

Build the binary

```sh
make
```

Start postgres

```sh
docker-compose up -Vd postgres
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1/operations\?sslmode=disable up
```

Run the web server

```sh
./bin/server
```

## Test

Start postgres

```sh
docker-compose up -Vd postgres
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1/operations\?sslmode=disable up
```

Run tests

```sh
go test -v ./...
```

## Usage

For example, you get information about this amazing Tesla Model X

```sh
http://localhost:8080/api/v1/operations?number=АА9359РС
```

```json
[
  {
    "number": "АА9359РС",
    "brand": "TESLA",
    "model": "MODEL X",
    "year": 2016,
    "date": "2016-10-13",
    "registration": "РЕЄСТРАЦIЯ ТЗ ПРИВЕЗЕНОГО З-ЗА КОРДОНУ ПО ВМД",
    "registration_code": 70,
    "fuel": "ЕЛЕКТРО",
    "capacity": null,
    "color": "ЧОРНИЙ",
    "kind": "ЛЕГКОВИЙ",
    "body": "УНІВЕРСАЛ-B",
    "purpose": "ЗАГАЛЬНИЙ",
    "own_weight": 2485,
    "total_weight": 3021,
    "reg_addr_koatuu": "8036600000",
    "dep_code": 8044,
    "dep": "Центр 8044",
    "person": "Юридична особа"
  }
]
```

## License

Project released under the terms of the MIT [license](./LICENSE).
