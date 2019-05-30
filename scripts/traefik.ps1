docker run --rm -d -p 8080:8080 -p 80:80 -v /var/run/docker.sock:/var/run/docker.sock traefik --api --docker
