docker run --rm -d -p 8080:8080 -p 80:80 traefik -v /var/run/docker.sock:/var/run/docker.sock --api --docker
