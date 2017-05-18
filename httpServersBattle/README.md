```bash
ab -c 1 -n 15000 http://127.0.0.1:7000/
ab -c 2 -n 7500 http://127.0.0.1:7000/
ab -c 50 -n 3000 http://127.0.0.1:7000/ (had to increase amount 10 times)
```
# Java
```
Requests per second:    2551.61 [#/sec] (mean)
Requests per second:    3512.87 [#/sec] (mean)
Requests per second:    3547.35 [#/sec] (mean)
```

# Go
# http
```
Requests per second:    3189.13 [#/sec] (mean)
Requests per second:    4927.39 [#/sec] (mean)
Requests per second:    8491.51 [#/sec] (mean)
```
# fasthttp
```
Requests per second:    4143.11 [#/sec] (mean)
Requests per second:    7490.37 [#/sec] (mean)
Requests per second:    8353.96 [#/sec] (mean)
```

# Python3.4
```
Requests per second:    2779.92 [#/sec] (mean)
Requests per second:    2390.44 [#/sec] (mean)
Can not handle
ab -c 50 -n 300 http://127.0.0.1:7000/ (10 times smaller)
Requests per second:    1427.92 [#/sec] (mean)
```