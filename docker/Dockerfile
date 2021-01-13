###############
### BUILDER ###
###############

FROM golang:1.15-alpine3.12 as builder

RUN apk add git build-base

WORKDIR /usr/src/microservice-email

ADD .git ./.git
ADD cmd ./cmd
ADD config ./config
ADD internal ./internal
ADD go.mod .
ADD go.sum .
ADD Makefile .
ADD LICENSE .

RUN make

###############
### RELEASE ###
###############

FROM alpine:3.12

LABEL Author="Sergio Andres Virviescas Santana <developersavsgio@gmail.com>"

COPY --from=builder /usr/src/microservice-email/ /microservice-email

RUN cd /microservice-email \
    && apk add --no-cache make git \
    && make install \
    && rm -rf /microservice-email \
    && apk del make git

# Configuration
COPY ./docker/docker-entrypoint.sh /usr/local/bin/
RUN ln -s /usr/local/bin/docker-entrypoint.sh /entrypoint.sh # backwards compat
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["microservice-email"]

EXPOSE 8000
