FROM alpine:3.14.1

MAINTAINER kazune-br

RUN apk update \
 && apk upgrade \
 && apk add --update-cache --no-cache --virtual .create-user shadow \
 && groupadd -g 1000 app \
 && useradd -u 1000 -g 1000 app \
 && apk del .create-user \
 && apk add ca-certificates gettext \
 && apk --update add tzdata \
 && mkdir -m 755 /app \
 && chown 1000:1000 /app \
 && rm -rf /var/cache/apk/*

USER app
WORKDIR /app

COPY ./.artifacts/app-linux-amd64 /app/
EXPOSE 8000