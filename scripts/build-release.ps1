param(
	[string] $version = "0.0.0"
)

$loc = Get-Location

cd ..

docker build -f webservice.dockerfile -t demo-webservice:$version --target web-service-release .
docker build -f authservice.dockerfile -t demo-authservice:$version --target auth-service-release .
docker build -f shopservice.dockerfile -t demo-shopservice:$version --target shop-service-release .

cd $loc
