FROM ubuntu:18.10 AS ubuntu-curl

ENV curl_version 7.54.1

RUN apt-get update -y && apt-get install -y build-essential wget

WORKDIR /tmp

RUN wget https://curl.haxx.se/download/curl-${curl_version}.tar.gz
RUN tar xfvz curl-${curl_version}.tar.gz
WORKDIR /tmp/curl-${curl_version}
RUN ./configure --disable-shared --enable-static --disable-threaded-resolver CFLAGS='-static -static-libgcc -Wl,-static -lc' && make && make install

##########################################################################################

FROM golang:1.12.4 AS shop-service-builder

RUN apt-get update && apt-get install -y git

COPY shopservice /src/shopservice
COPY go.mod /src
COPY go.sum /src
WORKDIR /src

ENV GO111MODULE=on
RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src/shopservice
RUN go build -ldflags="-w -s" -o shopservice

#########################################################################################

FROM golang:1.12.4 AS shop-service-dev

COPY --from=shop-service-builder /src/shopservice/shopservice /app/shopservice

EXPOSE 80

WORKDIR /app
ENTRYPOINT [ "/app/shopservice" ]

#########################################################################################

FROM scratch AS shop-service-release

COPY --from=shop-service-builder /src/shopservice/shopservice /shopservice
COPY --from=ubuntu-curl /usr/local/bin/curl /curl

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=30s --start-period=30s --retries=5 \
		CMD curl -f http://localhost/api/v1/health || exit 1

ENTRYPOINT [ "/shopservice" ]
