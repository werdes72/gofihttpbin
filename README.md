# gofihttpbin [![Go Report Card](https://goreportcard.com/badge/github.com/werdes72/gofihttpbin)](https://goreportcard.com/report/github.com/werdes72/gofihttpbin)

Fast port of httpbin written in Go using [Fiber](https://github.com/gofiber/fiber) web framework.
This project is still in development.

# Quickstart
```
docker run -p 80:8080 werd/gofihttpbin:main
curl -X GET "http://localhost/uuid"
```

# Performance:
gofihttpbin:
```
docker run --platform linux/amd64 -p 80:8080 werd/gofihttpbin:main
wrk -t 2 -c 4 -d 10s http://127.0.0.1/ip
Running 10s test @ http://127.0.0.1/ip
  2 threads and 4 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   594.68us    2.90ms  63.90ms   99.33%
    Req/Sec     5.04k   363.91     5.29k    98.02%
  101379 requests in 10.10s, 12.67MB read
Requests/sec:  10037.92
Transfer/sec:      1.25MB
```
httpbin:
```
docker run --platform linux/amd64 -p 80:80 kennethreitz/httpbin
wrk -t 2 -c 4 -d 10s http://127.0.0.1/ip
Running 10s test @ http://127.0.0.1/ip
  2 threads and 4 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.77ms    0.98ms  27.51ms   98.00%
    Req/Sec   297.61     14.86   313.00     93.00%
  5936 requests in 10.02s, 1.46MB read
Requests/sec:    592.29
Transfer/sec:    149.23KB
```

# Options
Enable request logging with an environment variable: `GOFI_LOGS=true`
