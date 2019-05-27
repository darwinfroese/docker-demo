param(
	[string]$version = 0.0.0
)

docker run --rm -d -l traefik.frontend.port=80 -l traefik.frontend.rule=Host:web.docker.demo --name demo-webservice webservice:$version
docker run --rm -d -l traefik.frontend.port=80 -l traefik.frontend.rule=Host:login.docker.demo --name demo-authservice authservice:$version
docker run --rm -d -l traefik.frontend.port=80 -l traefik.frontend.rule=Host:shop.docker.demo --name demo-shopservice shopservice:$version
