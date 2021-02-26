FROM golang:1.14.6-alpine

RUN apk update \
    && apk add --no-cache \
        autoconf \
        make \
        g++ \
        git \
        curl \
        openssl

ENV APP_USER app
ENV APP_GROUP app
ENV APP_HOME /go/src/app

RUN addgroup $APP_GROUP && adduser -G $APP_GROUP $APP_USER -D
RUN mkdir -p $APP_HOME && chown -R $APP_GROUP:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY ./main.go /go/src/app/main.go

RUN go get github.com/lib/pq
CMD ["go", "run", "main.go"]