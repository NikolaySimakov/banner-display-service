# Banner Display Service

[![GoDoc](https://godoc.org/github.com/lib/pq?status.svg)](https://pkg.go.dev/github.com/lib/pq?tab=doc)

Run project:

```bash
docker build -t my-golang-app .
docker run -it --rm --name my-running-app my-golang-app
```

## Introduction

The following technologies were used in the project:

- PostgreSQL (as data storage)
- pgx (driver for working with PostgreSQL)
- Docker (to start the service)
- Swagger (for API documentation)
- Echo (web framework)

### Life cicle

Create one feature:

POST: `http://localhost:8080/api/v1/features/`

```json
{
	"name": "90348094"
}
```

Create tags:

POST: `http://localhost:8080/api/v1/tags/`

```json
{
	"name": "34230489"
}

{
	"name": "04359903285"
}

{
	"name": "23094809832"
}
```

Create banner note:

POST: `http://localhost:8080/api/v1/banners/`

```json
{
	"title": "qwerty",
	"text": "we;kmewlkf;m lwf,m l,mw",
	"url": "wlkfmlwkemlmw",
	"feature": 0,
	"tag": [0, 1, 2],
	"active": true
}
```
