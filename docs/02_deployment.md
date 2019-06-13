# Deployment

## Docker

```bash
docker run -d --name bivac -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -e AWS_ACCESS_KEY_ID=XXXX \
  -e AWS_SECRET_ACCESS_KEY=XXXXXX \
  -e BIVAC_TARGET_URL=s3:my-bucket \
  -e RESTIC_PASSWORD=foo \
  -e BIVAC_SERVER_PSK=toto \
  -e BIVAC_AGENT_IMAGE=camptocamp/bivac:stable \
  -p 8182:8182 \
camptocamp/bivac:stable manager
```

You can easily deploy Bivac with a simple `docker-compose.yml`:

```yaml
version: '2'
services:
  bivac:
    image: camptocamp/bivac:stable
    command: manager
    ports:
      - "8182:8182"
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"
    environment:
      AWS_ACCESS_KEY_ID: XXXX
      AWS_SECRET_ACCESS_KEY: XXXXXX
      RESTIC_PASSWORD: foo
      BIVAC_TARGET_URL: s3:my-bucket
      BIVAC_SERVER_PSK: toto
      BIVAC_AGENT_IMAGE: camptocamp/bivac:stable
```

## Rancher (Cattle)

[Camptocamp](https://www.camptocamp.com) maintains a public template in its own [catalog](https://github.com/camptocamp/camptocamp-rancher-catalog).

As Rancher understands the docker-compose syntax, you can use the following example:

```yaml
version: '2'
services:
  bivac:
    image: camptocamp/bivac:stable
    command: manager
    ports:
      - "8182:8182"
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"
    environment:
      AWS_ACCESS_KEY_ID: XXXX
      AWS_SECRET_ACCESS_KEY: XXXXXX
      RESTIC_PASSWORD: foo
      BIVAC_TARGET_URL: s3:my-bucket
      BIVAC_SERVER_PSK: toto
      BIVAC_AGENT_IMAGE: camptocamp/bivac:stable
    labels:
      io.rancher.container.agent.role: environmentAdmin
      io.rancher.container.create_agent: 'true'
```

**Please note that you must run Bivac in a system stack.**

## Kubernetes
