# Nginx demo

A Nginx demo server as reverse proxy and load balancer including 7 different "hello world" apps, everything dockerized and orchestrated with docker-compose.

The structure of the base Nginx server is the following:

- Load balancer at `/`
  - Node.js (Express)
  - Python (Flask / Gunicorn)
  - PHP (Apache)
  - Go
  - Rust (Actix web)
- Apache server at `/apache` (serving static content)
- Nginx server at `/nginx` (serving static content)

They all return a "Hello world" message from their respective environments.

## Details

This is a personal learning project about Nginx, and therefore web servers, proxies and networking in general using Docker. The complexity that corresponds to the example applications is minimal since the effort is focused on the correct configuration and operation of a system of this nature, similar to a service-oriented or microservices architecture.

I've tried to optimize and follow best practices around images, using Alpine where possible, caching and other techniques. Still, this is probably not the best implementation since that's what it's all about... learning.

There are still challenges to overcome to deploy such a project but this is an excellent starting point to experiment and I am very happy with the result!

> You are free to explore and use the code at your convenience. I hope you find it useful and thanks for reading. ❤️
