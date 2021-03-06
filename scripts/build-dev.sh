#!/bin/bash

version=$1

if [ -z "$version" ]
then
	version="0.0.0"
fi

curr=`$pwd`
cd ..

docker build -f webservice.dockerfile -t demo-webservice:$version-dev --target web-service-dev .
docker build -f authservice.dockerfile -t demo-authservice:$version-dev --target auth-service-dev .
docker build -f shopservice.dockerfile -t demo-shopservice:$version-dev --target shop-service-dev . 

cd $curr
