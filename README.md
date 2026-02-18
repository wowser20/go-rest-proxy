# go-rest-proxy
A simple forward proxy server intended to secure client HTTP requests with JWT authentication proxies the https://dummyjson.com/ for the API responses.

I made this to familiarize the proxy implementation using REST API for the endpoints, as well as applying CLEAN architecture for the project structure. I haven't implemented dependency injection for this case, because I want this to be as simple as possible, since this is just a personal project for learning purposes.

### What is a proxy?
> A proxy, also called a forward proxy, routes traffic between clients and external systems, regulating it according to policies, masking IP addresses, enforcing security protocols, and blocking unknown traffic.

If you want to check the OpenAPI Documentation and test out the endpoints, you can visit the following URL: 
- https://staging-api-rest-proxy.aer-tech.me/docs/
(not available atm - just access it locally)

# Run locally

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

Then connect to this port to locally access the OpenAPI documentation

http://localhost:8080/docs



