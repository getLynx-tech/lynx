# lynx

Repository structure:
- the **client** folder contains the React Router frontend & backend.
- the **server** folder contains Go Gin backend.

## Local Development
You can run the project locally in a few steps using Docker Compose.

### SSL Certificates
Generate a self-signed ssh certificate using the following command and add it to your trust store.

```bash
openssl req -x509 -newkey rsa:4096 -sha256 -days 365 \
  -nodes -keyout .docker/nginx/ssl/dev.getlynx.tech.key -out .docker/nginx/ssl/dev.getlynx.tech.crt \
  -subj "/CN=dev.getlynx.tech" \
  -addext "subjectAltName=DNS:dev.getlynx.tech"
```

```bash
openssl req -x509 -newkey rsa:4096 -sha256 -days 365 \
  -nodes -keyout .docker/nginx/ssl/api.dev.getlynx.tech.key -out .docker/nginx/ssl/api.dev.getlynx.tech.crt \
  -subj "/CN=api.dev.getlynx.tech" \
  -addext "subjectAltName=DNS:api.dev.getlynx.tech"
```

Then add the development domain to your hosts file:
```bash
grep -qXF "127.0.0.1 api.dev.getlynx.tech dev.getlynx.tech" /etc/hosts || echo "127.0.0.1 api.dev.getlynx.tech dev.getlynx.tech" >> /etc/hosts
```

### Docker Compose
1. Navigate to the root directory of the project (where the `compose.yml` file is located).
2. Open a terminal in that directory.
3. Run the following command to build and start the Docker containers: `docker compose up --build -d`

This command will build the necessary Docker images and start the containers defined in the `compose.yml` file.
