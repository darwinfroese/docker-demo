#!/bin/bash

version=$1

docker run --rm -d -l traefik.frontend.port=80 -l traefik.frontend.rule=Host:web.docker.demo --name demo-webservice demo-webservice:$version
docker run --rm -d -l traefik.frontend.port=80 -l traefik.frontend.rule=Host:login.docker.demo --name demo-authservice demo-authservice:$version
docker run --rm -d -l traefik.frontend.port=80 -l traefik.frontend.rule=Host:shop.docker.demo --name demo-shopservice demo-shopservice:$version
