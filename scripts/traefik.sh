#!/bin/bash

docker run --rm -d -p 80:80 -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock traefik --api --docker
