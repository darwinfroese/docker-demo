FROM ubuntu:18.10 AS ubuntu-curl

ENV curl_version 7.54.1

RUN apt-get update -y && apt-get install -y build-essential wget

WORKDIR /tmp

RUN wget https://curl.haxx.se/download/curl-${curl_version}.tar.gz
RUN tar xfvz curl-${curl_version}.tar.gz
WORKDIR /tmp/curl-${curl_version}
RUN ./configure --disable-shared --enable-static --disable-threaded-resolver CFLAGS='-static -static-libgcc -Wl,-static -lc' && make && make install

##########################################################################################

FROM golang:1.12.4 AS auth-service-builder

RUN apt-get update && apt-get install -y git

COPY authservice /src/authservice
COPY go.mod /src
COPY go.sum /src
WORKDIR /src

ENV GO111MODULE=on
RUN go mod download

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src/authservice
RUN go build -ldflags="-w -s" -o authservice

#########################################################################################

FROM golang:1.12.4 AS auth-service-dev

COPY --from=auth-service-builder /src/authservice/authservice /app/authservice

EXPOSE 80

WORKDIR /app
ENTRYPOINT [ "/app/authservice" ]

#########################################################################################

FROM scratch AS auth-service-release

COPY --from=auth-service-builder /src/authservice/authservice /authservice
COPY --from=ubuntu-curl /usr/local/bin/curl /curl

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=30s --start-period=30s --retries=5 \
		CMD curl -f http://localhost/api/v1/health || exit 1

ENTRYPOINT [ "/authservice" ]
