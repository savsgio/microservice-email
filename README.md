# microservice-email

[![Build Status](https://travis-ci.org/savsgio/microservice-email.svg?branch=master)](https://travis-ci.org/savsgio/microservice-email)
[![Go Report Card](https://goreportcard.com/badge/github.com/savsgio/microservice-email)](https://goreportcard.com/report/github.com/savsgio/microservice-email)
[![GitHub release](https://img.shields.io/github/release/savsgio/microservice-email.svg)](https://github.com/savsgio/microservice-email/releases)

Microservice to send emails

### System requirements

- [Go](https://golang.org/dl/) (>= 1.11)
- [RabbitMQ](https://www.rabbitmq.com/)
- make

## Installation

Download Go dependencies and build:

```bash
make
```

Install

```bash
make install
```

After, you can exec with

```bash
microservice-email
```

Optional arguments:

- `-log-level`: Level of log that you want to show (default: _info_)
- `-config-file`: Path of configuration file (default: _/etc/microservice-email.yml_)
- `-version`: Print version of service

### API:

This API only accept **_POST_** http request with below parameters in body:

Explanation (all are required):

- `to`: List of emails of destiny
- `subject`: Subject of email
- `content_type`: Content type of email that it can be **_text/plain_** or **_text/html_**
- `body`: Content of email

Example of request to send a email:

```json
{
  "to": ["example_1@example.com", "example_2@example.com"],
  "subject": "Hi, my friend",
  "content_type": "text/html",
  "body": "<h1>This is the body of my Email in HTML format</h1>"
}
```

## Docker

### Dependencies

- [Docker](https://www.docker.com/)
- [Docker-compose](https://docs.docker.com/compose/) \_Recommended to install with `pip3` (python3).

Build:

```bash
make docker_build
```

Run:

```bash
make docker_run
```

## For Devs

Copy `config/microservice-email.conf` to `config/microservice-email.dev.conf.yml` _(this file not tracked in Git)_, modify each config and exec:

```bash
make run
```

**_Note:_** If you want to use with Docker, make sure you have this rabbitmq configuration in `config/microservice-email.dev.conf.yml`:

```yaml
...
rabbitmq:
  host: rabbitmq_server
  user: guest
  password: guest
  ...
...
```

## Contributing

**Feel free to contribute it or fork me...** :wink:
