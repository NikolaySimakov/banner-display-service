# Banner Display Service

[![GoDoc](https://godoc.org/github.com/lib/pq?status.svg)](https://pkg.go.dev/github.com/lib/pq?tab=doc)

Run project:

```bash
docker build -t banner-display-service .
docker run -it --rm --name run-banner-display-service banner-display-service
```

## Introduction

What I added besides the specified handlers in the `api.yaml`:

- auth router
- tags router
- features router

The following technologies were used in the project:

- PostgreSQL (as data storage)
- pgx (driver for working with PostgreSQL)
- Docker (to start the service)
- Swagger (for API documentation)
- Echo (web framework)

### Usage

First, you should log in:

![1.png](/docs/images/3.png)

#### Create features

POST: `http://localhost:8080/api/v1/features/`

```curl
curl -X POST http://localhost:8080/api/v1/features/ \
     -H "Content-Type: application/json" \
     -H "token: 2cd38435c3e08658a18ad97b3f1d83419857955fce336d7d868098652e0957f5" \
		 -d '{"name": "123456789"}'
```

```json
{
	"name": "123456789"
}
```

#### Create tags

POST: `http://localhost:8080/api/v1/tags/`

![3.png](/docs/images/3.png)

```curl
curl -X POST http://localhost:8080/api/v1/tags/ \
     -H "Content-Type: application/json" \
     -H "token: 2cd38435c3e08658a18ad97b3f1d83419857955fce336d7d868098652e0957f5" \
		 -d '{"name": "123456789"}'
```

Response:

```json
{
	"name": "123456789"
}
```

#### Create banner

POST: `http://localhost:8080/api/v1/banners/`

```curl
curl -X POST http://localhost:8080/api/v1/banners/ \
     -H "Content-Type: application/json" \
     -H "token: 2cd38435c3e08658a18ad97b3f1d83419857955fce336d7d868098652e0957f5" \
     -d '{"title": "qwerty", "text": "dsljfwhjfljhwf", "url": "wlkfmlwkemlmw", "feature_id": 111, "tag_id": [123, 234, 456], "is_active": true}'
```

Response:

```json
{
	"title": "qwerty",
	"text": "dsljfwhjfljhwf",
	"url": "wlkfmlwkemlmw",
	"feature_id": 111,
	"tag_id": [123, 234, 456],
	"is_active": true
}
```

#### Get banners for user (use user token)

GET: `http://localhost:8080/api/v1/banners/user_banner`

Query params:

- tag_id
- feature_id
- use_last_revision

For case: GET `http://localhost:8080/api/v1/banners/user_banner?tag_id=2&feature_id=2&use_last_revision=true`

```curl
curl -X GET http://localhost:8080/api/v1/banners/user_banner?tag_id=2&feature_id=2&use_last_revision=true \
-H "token: 55f6eae4fd2bf89a63ce2cefa79e351d8c768fb51820d7b5fd9e246d7f0ffff8"
```

![2.png](/docs/images/2.png)

Response:

```json
[
	{
		"Id": 16,
		"Title": "wejhggehw",
		"Text": "wwepjwfpj",
		"Url": "wlekfhjlkjfhw",
		"CreatedAt": "2024-04-13T20:32:27.866533Z",
		"UpdatedAt": {
			"Time": "0001-01-01T00:00:00Z",
			"Valid": false
		},
		"LastVersion": true,
		"IsActive": true,
		"TagId": [2, 3],
		"FeatureId": 2
	}
]
```

#### Get banners for admin (use admin token)

GET `http://localhost:8080/api/v1/banners/banner`

Query params:

- tag_id
- feature_id
- limit
- offset

```curl
curl -X GET http://localhost:8080/api/v1/banners/banner \
     -H "token: 2cd38435c3e08658a18ad97b3f1d83419857955fce336d7d868098652e0957f5"
```

![4.png](/docs/images/4.png)

```json
[
	{
		"Id": 16,
		"Title": "wejhggehw",
		"Text": "wwepjwfpj",
		"Url": "wlekfhjlkjfhw",
		"CreatedAt": "2024-04-13T20:32:27.866533Z",
		"UpdatedAt": {
			"Time": "0001-01-01T00:00:00Z",
			"Valid": false
		},
		"LastVersion": true,
		"IsActive": true,
		"TagId": [2, 3],
		"FeatureId": 2
	},
	{
		"Id": 17,
		"Title": "eroigjoreigjk",
		"Text": "gkhrtjlgh",
		"Url": "dlkjkdfwkj",
		"CreatedAt": "2024-04-13T20:58:07.375875Z",
		"UpdatedAt": {
			"Time": "0001-01-01T00:00:00Z",
			"Valid": false
		},
		"LastVersion": true,
		"IsActive": true,
		"TagId": [2, 3, 4],
		"FeatureId": 1
	},
	{
		"Id": 15,
		"Title": "qwerty",
		"Text": "we;kmewlkf;m lwf,m l,mw",
		"Url": "wlkfmlwkemlmw",
		"CreatedAt": "2024-04-13T20:27:09.189635Z",
		"UpdatedAt": {
			"Time": "0001-01-01T00:00:00Z",
			"Valid": false
		},
		"LastVersion": true,
		"IsActive": false,
		"TagId": [2, 3, 4],
		"FeatureId": 2
	}
]
```

#### Delete banner

DELETE `http://localhost:8080/api/v1/banners/`

```curl
curl -X DELETE 'http://localhost:8080/api/v1/banners/' \
     -H 'Content-Type: application/json' \
     -H 'token: 2cd38435c3e08658a18ad97b3f1d83419857955fce336d7d868098652e0957f5' \
     -d '{"feature_id": 2, "tag_id": 2}'
```

Body:

```json
{
	"feature_id": 2,
	"tag_id": 2
}
```

![5.png](/docs/images/5.png)

Response:

No content 204
