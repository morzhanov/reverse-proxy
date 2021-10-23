# reverse-proxy

Reverse proxy solutions example. 

Repo contains basic reverse proxy examples with popular solutions:

- Apache Server
- Envoy
- HAProxy
- NGINX

Example based on two identical servers which are load balanced by proxies.

## Run

```bash
docker-compose -p rproxy up -d
```

- Apache Server: http:127.0.0.1:3001
- Envoy: http:127.0.0.1:3002
- HAProxy: http:127.0.0.1:3003
- NGINX: http:127.0.0.1:3004
