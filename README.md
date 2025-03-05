# go-rest-proxy
A simple forward proxy server intended for securing HTTP requests with JWT authorization, utilizes https://dummyjson.com/ for the dummy API responses.

> What is a proxy?
> A proxy server routes traffic between client(s) and external system, regulating it according to policies, masking IP addresses, enforcing security protocols, and blocking unknown traffic.

# How to run locally

Initialize the environment variables, .env file

```bash
cp .env.example .env
```

To build the application or check if it's working, run:

```bash
make build
```

To run or start the proxy application: build, install, and run:

```bash
make 
```

To install, run:

``` bash
make install
```

To test, run:

```bash
make test
```


TBD - URL for OpenAPI docs



