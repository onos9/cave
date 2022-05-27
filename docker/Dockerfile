
###############################
# STEP 2 Build services image
###############################
FROM golang:1.17-alpine AS builder

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GOARCH=amd64

WORKDIR /app

# System dependencies
RUN apk update && apk upgrade \
    && apk add --no-cache \
    git nodejs npm ca-certificates

COPY go.* /app/

# Fetch dependencies.
RUN  go mod tidy -compat=1.17 \
    && go mod download \
    && go mod verify

COPY  . /app/

# Buid for production
RUN cd cmd/api \
    && go build -gcflags "all=-N -l" -o ../../api

RUN git clone https://github.com/onos9/cave-landing.git web
RUN git clone https://github.com/onos9/cave-web.git portal \
    && cd portal \
    && npm install --legacy-peer-deps \
    && npm run build

# COPY cmd/api/templates templates
# ENTRYPOINT ["./api"]

################################
# STEP 3 build a small image
################################
FROM scratch AS api
COPY --from=builder /app/api ./api
COPY --from=builder /app/cmd/api/templates /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./api"]

################################
# STEP 4 build for nginx
################################
FROM nginx:alpine AS nginx
COPY docker/nginx/default.conf /etc/nginx/conf.d
COPY --from=builder /app/portal/build/ /var/www/html/portal
COPY --from=builder /app/web/ /var/www/html/web

################################
# STEP 5 build for php
################################
FROM php:fpm-alpine As php
RUN docker-php-ext-install opcache
COPY docker/php/opcache.ini /usr/local/etc/php/conf.d/opcache.ini
COPY --from=builder /app/web/ /var/www/html/web