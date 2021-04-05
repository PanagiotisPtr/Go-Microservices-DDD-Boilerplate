# Basic Go Microservice setup

## Overview
This repo contains some boilerplate code to create a very basic Go microservice using docker and docker-compose. It uses Go-kit, PostgreSQL and PgAdmin.

## To configure microservice
There are 2 places where config files exist. One is under `account-service/config/*.yml` and one is in `config/*.env`. The `./config` directory contains all the configuration needed by docker and the related services and `./account-service/config` has config files related to the service itself. There are example configs in each directory called `docker.dist.env` and `service.dist.yml` which you can use to set up your own configs. You will need 4 new configs in total (2 for prod and 2 for dev). Those should be:

```
config/docker.dev.env
config/docker.prod.env
account-service/config/service.dev.yml
account-service/config/service.prod.yml
```

## To start microservice
Run
```sh
docker-compose --env-file ./config/docker.{env}.env -f stack.{env}.yml up --build
```
Where `{env}` is either `dev` for development or `prod` for production. When building for production, the pgadmin service isn't built. Feel free to add `-d` to the end to run in detached mode.

## To stop microservice
If you want to remove the previous containers (this *does not* remove the volumes - aka doesn't re-initialise the database) run

```sh
docker-compose --env-file ./config/docker.{env}.env -f stack.{env}.yml down
```