FROM ubuntu:18.10 AS ubuntu-curl

ENV curl_version 7.54.1

RUN apt-get update -y && apt-get install -y build-essential wget

WORKDIR /tmp

RUN wget https://curl.haxx.se/download/curl-${curl_version}.tar.gz
RUN tar xfvz curl-${curl_version}.tar.gz
WORKDIR /tmp/curl-${curl_version}
RUN ./configure --disable-shared --enable-static --disable-threaded-resolver CFLAGS='-static -static-libgcc -Wl,-static -lc' && make && make install

##########################################################################################

FROM golang:1.12.4 AS web-service-builder

RUN apt-get update && apt-get install -y git

COPY webservice /src/webservice
COPY go.mod /src
COPY go.sum /src
WORKDIR /src

ENV GO111MODULE=on
RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src/webservice
RUN go build -ldflags="-w -s" -o webservice

#########################################################################################

FROM node:12.2.0 AS web-builder

RUN mkdir /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node-modules/.bin:$PATH

COPY webservice-frontend/package.json /usr/src/app/package.json
RUN npm install --silent
RUN npm install react-scripts@2.1.3 -g --silent

COPY webservice-frontend/src /usr/src/app/src
COPY webservice-frontend/public /usr/src/app/public

RUN npm run-script build --silent

#########################################################################################

FROM golang:1.12.4 AS web-service-dev

COPY --from=web-service-builder /src/webservice/webservice /app/webservice
COPY --from=web-builder /usr/src/app/build /app/www

EXPOSE 80

WORKDIR /app
ENTRYPOINT [ "/app/webservice" ]

#########################################################################################

FROM scratch AS web-service-release

COPY --from=web-service-builder /src/webservice/webservice /webservice
COPY --from=ubuntu-curl /usr/local/bin/curl /curl
COPY --from=web-builder /usr/src/app/build /www

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=30s --start-period=30s --retries=5 \
		CMD curl -f http://localhost/api/v1/health || exit 1

ENTRYPOINT [ "/webservice" ]
