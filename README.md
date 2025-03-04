# go-rest-proxy
- A simple forward proxy intended for securing requests equipped with jwt authorization, utilizes https://dummyjson.com/ for the dummy API responses.

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



