# go-rest-proxy
A simple forward proxy server intended to secure client HTTP requests with JWT authentication utilizes https://dummyjson.com/ for the dummy API responses.

I made this to familiarize the proxy implementation using REST API for the endpoints, as well as Domain Driven Design Pattern for the project structure, I haven't implemented dependency injection for this cause I want this to be as simple as possible since this is just a personal project for learning purposes.

### What is a proxy?
> A proxy, also called a forward proxy, routes traffic between clients and external system, regulating it according to policies, masking IP addresses, enforcing security protocols, and blocking unknown traffic.



If you want to check the OpenAPI Documentation and test out the endpoints, you can visit the following URL: 
- https://staging-api-rest-proxy.aer-tech.org/docs/
(make sure to use the staging url in the servers dropdown in the documentation)

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



