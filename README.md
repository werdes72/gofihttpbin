# gofihttpbin [![Go Report Card](https://goreportcard.com/badge/github.com/werdes72/gofihttpbin)](https://goreportcard.com/report/github.com/werdes72/gofihttpbin)

Fast port of httpbin written in Go using [Fiber](https://github.com/gofiber/fiber) web framework.
This project is still in development.

# Quickstart
```
docker run -p 80:8080 werd/gofihttpbin:main
curl -X GET "http://localhost/uuid"
```
