# chibi_linku
URL shortener in Golang

## Routes

```
GET /ping // ping the service
GET /purge // purge the content of the redis
POST /encode // encode an URL and return its identifier
GET /decode/{code} // decode an identifier and redirect to the related url
```
