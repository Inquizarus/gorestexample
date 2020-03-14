# GOREST example
This is just a quick example server using my [REST library](https://github.com/inquizarus/gorest).

Start server with
```bash
$ go run main.go
```

Then test it with for example curl
```bash
$ curl localhost:8080
...
{"message": "Hello, World!"}
```

If you want to test out TLS connections you can run `generate_certs` (requires openssl) to
create a self signed one. Enable TLS by changing line `12` in main.go and
set the `Enable` field to `true` and restart the server.

Run curl against it, now with HTTPS protocol. (-k is mainly for ignoring that the certificate is self signed)
```bash
$ curl -k https://localhost:8080
...
{"message": "Hello, World!"}
```

There are other settings to play around with, feel free to do so. :)
