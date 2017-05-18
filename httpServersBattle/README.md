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
```
<details>
  <summary>Could not handle</summary>
  
```python
  Exception happened during processing of request from ('127.0.0.1', 53634)
  Traceback (most recent call last):
    File "/usr/lib/python3.4/socketserver.py", line 305, in _handle_request_noblock
      self.process_request(request, client_address)
    File "/usr/lib/python3.4/socketserver.py", line 331, in process_request
      self.finish_request(request, client_address)
    File "/usr/lib/python3.4/socketserver.py", line 344, in finish_request
      self.RequestHandlerClass(request, client_address, self)
    File "/usr/lib/python3.4/socketserver.py", line 669, in __init__
      self.handle()
    File "/usr/lib/python3.4/http/server.py", line 398, in handle
      self.handle_one_request()
    File "/usr/lib/python3.4/http/server.py", line 386, in handle_one_request
      method()
    File "httpServer.py", line 18, in do_GET
      self.wfile.write(bytes(message, "utf8"))
    File "/usr/lib/python3.4/socket.py", line 391, in write
      return self._sock.send(b)
  BrokenPipeError: [Errno 32] Broken pipe
```
</details>

```
ab -c 50 -n 300 http://127.0.0.1:7000/ (10 times less)
Requests per second:    1427.92 [#/sec] (mean)
```