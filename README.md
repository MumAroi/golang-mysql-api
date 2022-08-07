## Installation

```
docker-compose up -d
```

Or:

```
docker-compose up --build
```

## Rebuild

```
docker-compose down --remove-orphans --volumes
docker-compose up --build
```


## Get login
### Request

`GET /login`
    curl -i -H 'Accept: application/json' http://localhost:7000/thing/
### Response

    HTTP/1.1 201 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /thing/1
    Content-Length: 36

    {"token": "tokenhere"}